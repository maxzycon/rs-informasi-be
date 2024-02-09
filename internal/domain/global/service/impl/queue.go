package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/constant/status"
	typequeuedetail "github.com/maxzycon/rs-farmasi-be/pkg/constant/type_queue_detail"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetQueuePaginated(ctx context.Context, payload *dto.ParamsQueueQueries) (resp dto.QueueWrapper, err error) {
	resp = dto.QueueWrapper{
		Summary: &dto.SummaryQueue{
			Total:           0,
			TotalDiserahkan: 0,
			TotalRacikan:    0,
			TotalNonRacikan: 0,
		},
		Items: &dto.QueueDataWrapper{
			Queues: make([]*dto.QueueRowPaginated, 0),
			Paginator: dto.DefaultPaginationDtoRow{
				CurrentPage:   payload.Page,
				RecordPerPage: payload.Limit,
				LastPage:      0,
				TotalItem:     0,
			},
		},
	}

	if payload.Limit < 1 {
		payload.Limit = 1
	}

	payload.Page = payload.Page - 1
	user, _ := authutil.GetCredential(ctx)

	cond := squirrel.And{
		squirrel.Eq{
			"q.deleted_at": nil,
		},
	}

	if user.MerchantID != nil {
		cond = append(cond, squirrel.Eq{
			"q.merchant_id": user.MerchantID,
		})
	}

	if len(payload.LocationIDS) > 0 {
		cond = append(cond, squirrel.Eq{
			"q.location_id": payload.LocationIDS,
		})
	}

	if len(payload.Type) > 0 {
		cond = append(cond, squirrel.Eq{
			"q.type": payload.Type,
		})
	}

	if payload.Search != nil && *payload.Search != "" {
		cond = append(cond, squirrel.Or{
			squirrel.Like{
				"q.medical_record": fmt.Sprintf("%%%s%%", *payload.Search),
			},
			squirrel.Like{
				"q.queue_no": fmt.Sprintf("%%%s%%", *payload.Search),
			},
		})
	}

	if len(payload.Status) > 0 {
		cond = append(cond, squirrel.Eq{
			"l.status": payload.Status,
		})
	}

	getStr, argStr, err := squirrel.Select(`q.id, q.id_str, q.queue_no, q.medical_record, q.location_id, q.location_name, q.type, q.merchant_id, q.merchant_name, q.user_id, q.user_name, l.status as lastStatus, COALESCE(lastExtend.end_queue, process.end_queue) as endTime, q.created_at`).
		From("queues as q").
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as process ON process.queue_id = q.id AND process.status = 2").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as lastExtend ON lastExtend.queue_id = q.id AND lastExtend.status = 2 AND lastExtend.type = 2").
		Where(cond).
		Limit(payload.Limit).
		Offset(payload.Limit * payload.Page).
		OrderBy(fmt.Sprintf("q.%s %s", payload.SortBy, payload.Order)).
		ToSql()

	if err != nil {
		return
	}

	row, err := s.db.WithContext(ctx).Raw(getStr, argStr...).Rows()
	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.QueueRowPaginated{}
		err = row.Scan(
			&temp.ID, &temp.IDStr, &temp.QueueNo, &temp.MedicalRecord,
			&temp.LocationID, &temp.LocationName, &temp.Type, &temp.MerchantID,
			&temp.MerchantName, &temp.CreatedByID, &temp.CreatedBy, &temp.Status,
			&temp.EstEnd, &temp.CreatedAt,
		)
		if err != nil {
			return
		}

		resp.Items.Queues = append(resp.Items.Queues, &temp)
	}

	// ----- Count Pagination
	getStrCount, argStrCount, err := squirrel.Select(`COUNT(q.id) as id`).
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as process ON process.queue_id = q.id AND process.status = 2").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as lastExtend ON lastExtend.queue_id = q.id AND lastExtend.status = 2 AND lastExtend.type = 2").
		From("queues q").
		Where(cond).
		ToSql()
	if err != nil {
		return
	}

	var totalRows int64
	err = s.db.WithContext(ctx).Raw(getStrCount, argStrCount...).Row().Scan(&totalRows)
	if err != nil {
		return
	}

	total := totalRows / int64(payload.Limit)
	remainder := totalRows % int64(payload.Limit)

	if remainder == 0 {
		resp.Items.Paginator.LastPage = uint64(total)
	} else {
		resp.Items.Paginator.LastPage = uint64(total + 1)
	}
	resp.Items.Paginator.RecordPerPage = payload.Limit
	resp.Items.Paginator.TotalItem = uint64(totalRows)
	// ----- End count pagination

	// ----- Summary total
	err = s.db.Raw(`SELECT COALESCE(COUNT(id),0) as total FROM queues WHERE CAST(created_at as DATE) = CAST(now() as DATE) AND deleted_at IS NULL`).Scan(&resp.Summary.Total).Error
	if err != nil {
		return
	}

	// ----- Summary total diserahkan
	err = s.db.Raw(`SELECT COALESCE(COUNT(id),0) as total
	FROM queues q
	LEFT JOIN (SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as l ON l.queue_id = q.id
	WHERE CAST(created_at as DATE) = CAST(now() as DATE) 
	AND deleted_at IS NULL
	AND l.status = 4`).Scan(&resp.Summary.TotalDiserahkan).Error
	if err != nil {
		return
	}

	// ----- Summary racikan
	err = s.db.Raw(`SELECT COALESCE(COUNT(id),0) as total FROM queues WHERE type = 2 AND CAST(created_at as DATE) = CAST(now() as DATE) AND deleted_at IS NULL`).Scan(&resp.Summary.TotalRacikan).Error
	if err != nil {
		return
	}

	// ----- Summary non racikan
	err = s.db.Raw(`SELECT COALESCE(COUNT(id),0) as total FROM queues WHERE type = 1 AND CAST(created_at as DATE) = CAST(now() as DATE) AND deleted_at IS NULL`).Scan(&resp.Summary.TotalNonRacikan).Error
	if err != nil {
		return
	}

	return
}

// func (s *GlobalService) GetQueueById(ctx context.Context, id int) (resp *dto.QueueRow, err error) {
// 	row, err := s.globalRepository.FindQueueById(ctx, id)
// 	if err != nil {
// 		log.Errorf("err get Queue paginated")
// 		return
// 	}
// 	resp = &dto.QueueRow{
// 		ID:   row.ID,
// 		Name: row.Name,
// 	}
// 	return
// }

func (s *GlobalService) CreateQueue(ctx context.Context, payload *dto.PayloadQueue) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	err = s.db.Transaction(func(tx *gorm.DB) error {

		queue := &model.Queue{
			LocationID:    payload.LocationID,
			QueueNo:       payload.QueueNo,
			MedicalRecord: payload.MedicalRecord,
			Type:          payload.Type,
			UserID:        user.ID,
			Histories: []model.QueueHistory{
				{
					Status:     status.VALIDATION,
					StartQueue: nil,
					Duration:   nil,
					EndQueue:   nil,
					Type:       typequeuedetail.DEFAULT,
					UserID:     user.ID,
					UserName:   user.Name,
				},
			},
			UserName: user.Name,
		}

		if payload.MerchantID != nil {
			queue.MerchantID = *payload.MerchantID
		}

		if user.MerchantID != nil {
			queue.MerchantID = *user.MerchantID
		}

		// ----- get merchant
		merchant, err := s.globalRepository.FindMerchantById(ctx, int(queue.MerchantID))
		if err != nil {
			return err
		}
		queue.MerchantName = merchant.Name

		// ----- get location
		location, err := s.globalRepository.FindLocationById(ctx, int(queue.MerchantID))
		if err != nil {
			return err
		}
		queue.LocationName = location.Name

		if err := tx.Create(queue).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return
	}
	success := int64(1)
	resp = &success
	return
}

func (s *GlobalService) UpdateStatusQueueById(ctx context.Context, id int, payload *dto.PayloadUpdateQueue) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	err = s.db.Transaction(func(tx *gorm.DB) error {

		queue := &model.QueueHistory{
			Status:     payload.NewStatus,
			StartQueue: nil,
			Duration:   nil,
			EndQueue:   nil,
			Type:       typequeuedetail.DEFAULT,
			UserID:     user.ID,
			QueueID:    uint(id),
			UserName:   user.Name,
		}

		// ----- process default / extend
		if payload.NewStatus == status.PROCESS {
			now := time.Now().Local()
			end := now.Add(time.Hour +
				time.Minute*time.Duration(payload.Duration) +
				time.Second)
			queue.StartQueue = &now
			queue.Duration = &payload.Duration
			queue.EndQueue = &end
		}

		if err := tx.Create(queue).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return
	}
	success := int64(1)
	resp = &success
	return
}

func (s *GlobalService) DeleteQueueById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteQueueById(ctx, id)
	if err != nil {
		log.Errorf("err delete Queue %d", id)
		return
	}
	return
}

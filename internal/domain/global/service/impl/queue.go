package impl

import (
	"context"
	"database/sql"
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

	getStr, argStr, err := squirrel.Select(`q.id, q.id_str, q.queue_no, q.medical_record, q.location_id, q.location_name, q.type, q.merchant_id, q.merchant_name, q.user_id, q.user_name, l.status as lastStatus, COALESCE(lastExtend.end_queue, process.end_queue) as endTime, q.created_at, q.is_follow_up, q.follow_up_phone, 
	(CASE WHEN lastExtend.end_queue IS NOT NULL THEN 1 ELSE 0 END) as isExtend`).
		From("queues as q").
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2)) as process ON process.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2 AND qh.type = 2)) as lastExtend ON lastExtend.queue_id = q.id").
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
			&temp.EstEnd, &temp.CreatedAt, &temp.IsFollowUp, &temp.FollowUpPhone,
			&temp.IsExtend,
		)
		if err != nil {
			return
		}

		resp.Items.Queues = append(resp.Items.Queues, &temp)
	}

	// ----- Count Pagination
	getStrCount, argStrCount, err := squirrel.Select(`COUNT(q.id) as id`).
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id)) as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2)) as process ON process.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2 AND qh.type = 2)) as lastExtend ON lastExtend.queue_id = q.id").
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

func (s *GlobalService) UpdateFuQueueNo(ctx context.Context, id string, newPhone string) (err error) {
	update, args, err := squirrel.
		Update("queues").
		Set("is_follow_up", 1).
		Set("follow_up_phone", newPhone).
		Where(squirrel.Eq{
			"id_str": id,
		}).
		ToSql()

	if err != nil {
		log.Error("err sql update")
		return
	}

	err = s.db.Exec(update, args...).Error

	if err != nil {
		log.Error("err update")
		return
	}

	return
}

func (s *GlobalService) GetDashboardDisplay(ctx context.Context, payload *dto.ParamsQueueDisplay, merchantIdStr string) (resp *dto.QueueDataDisplayWrapper, err error) {
	resp = &dto.QueueDataDisplayWrapper{
		Queues: make([]*dto.QueueDataDisplay, 0),
		Paginator: dto.DefaultPaginationDtoRow{
			CurrentPage:   payload.Page,
			RecordPerPage: payload.Limit,
			LastPage:      0,
			TotalItem:     0,
		},
	}
	user, _ := authutil.GetCredential(ctx)

	if payload.Limit < 1 {
		payload.Limit = 1
	}

	payload.Page = payload.Page - 1

	cond := squirrel.And{
		squirrel.Eq{
			"q.deleted_at": nil,
		},
		squirrel.LtOrEq{
			"COALESCE(TIMESTAMPDIFF(MINUTE, diterima.created_at, NOW()),0)": 5, // 5 menit
		},
	}

	if user.MerchantID != nil {
		cond = append(cond, squirrel.Eq{
			"q.merchant_id": *user.MerchantID,
		})
	}

	if merchantIdStr != "" {
		cond = append(cond, squirrel.Eq{
			"m.id_str": payload.MerchantIdStr,
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

	// ----- Data
	getParentStr, argsParent, err := squirrel.
		Select("q.id, q.medical_record, q.queue_no, q.type, l.status, COALESCE(lastExtend.end_queue, process.end_queue, '-') as endTime").
		From("queues as q").
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2)) as process ON process.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2 AND qh.type = 2)) as lastExtend ON lastExtend.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 4)) as diterima ON diterima.queue_id = q.id").
		InnerJoin("merchants m ON m.id = q.merchant_id").
		Where(cond).
		Limit(payload.Limit).
		Offset(payload.Limit * payload.Page).
		OrderBy("l.status ASC, q.created_at ASC").
		ToSql()

	if err != nil {

		return
	}

	row, err := s.db.WithContext(ctx).Raw(getParentStr, argsParent...).Rows()
	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.QueueDataDisplay{}
		err = row.Scan(
			&temp.ID, &temp.MedicalRecord, &temp.QueueNo,
			&temp.Type, &temp.Status, &temp.EstEnd,
		)
		if err != nil {
			return
		}

		resp.Queues = append(resp.Queues, &temp)
	}

	// ----- End Data

	// ----- Count Pagination
	getStrCount, argStrCount, err := squirrel.Select(`COUNT(q.id) as id`).
		From("queues as q").
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2)) as process ON process.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2 AND qh.type = 2)) as lastExtend ON lastExtend.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 4)) as diterima ON diterima.queue_id = q.id").
		InnerJoin("merchants m ON m.id = q.merchant_id").
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
		resp.Paginator.LastPage = uint64(total)
	} else {
		resp.Paginator.LastPage = uint64(total + 1)
	}
	resp.Paginator.RecordPerPage = payload.Limit
	resp.Paginator.TotalItem = uint64(totalRows)
	// ----- End count pagination

	return
}

func (s *GlobalService) GetQueueBySearch(ctx context.Context, merchantId string, search string) (resp *dto.QueueUserSearch, err error) {
	resp = &dto.QueueUserSearch{}
	cond := squirrel.And{
		squirrel.Eq{
			"q.deleted_at": nil,
		},
		squirrel.Eq{
			"m.id_str": merchantId,
		},
	}

	if search != "" {
		cond = append(cond, squirrel.Or{
			squirrel.Like{
				"q.medical_record": fmt.Sprintf("%%%s%%", search),
			},
			squirrel.Like{
				"q.queue_no": fmt.Sprintf("%%%s%%", search),
			},
		})
	}

	// ----- Data
	getParentStr, argsParent, err := squirrel.
		Select("q.id_str, q.medical_record, q.queue_no, q.type, l.status, COALESCE(lastExtend.end_queue, process.end_queue, '-') as endTime, q.is_follow_up, q.follow_up_phone").
		From("queues as q").
		LeftJoin("(SELECT queue_id, status FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id))  as l ON l.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2)) as process ON process.queue_id = q.id").
		LeftJoin("(SELECT * FROM queue_histories as c WHERE id = (SELECT MAX(qh.id) FROM queue_histories as qh WHERE qh.queue_id = c.queue_id AND qh.status = 2 AND qh.type = 2)) as lastExtend ON lastExtend.queue_id = q.id").
		InnerJoin("merchants m ON m.id = q.merchant_id").
		Where(cond).
		Limit(1).
		OrderBy("q.id DESC").
		ToSql()

	if err != nil {
		return
	}

	err = s.db.Raw(getParentStr, argsParent...).
		Row().
		Scan(&resp.ID, &resp.MedicalRecord, &resp.QueueNo,
			&resp.Type, &resp.Status, &resp.EstEnd, &resp.IsFollowUp, &resp.FollowUpPhone)

	if err == sql.ErrNoRows {
		err = gorm.ErrRecordNotFound
		return
	}

	if err != nil {
		return
	}

	return
}

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
		location, err := s.globalRepository.FindLocationById(ctx, int(payload.LocationID))
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

		if payload.Type == typequeuedetail.EXTEND {
			queue.Type = typequeuedetail.EXTEND
		}

		// ----- process default / extend
		if payload.NewStatus == status.PROCESS {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			now := time.Now().In(loc)
			end := now.Add(time.Minute * time.Duration(payload.Duration))
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

func (s *GlobalService) GetQueueById(ctx context.Context, id int) (resp *dto.QueueRowDetail, err error) {
	data, err := s.globalRepository.FindQueueById(ctx, id)
	if err != nil {
		log.Errorf("err delete Queue %d", id)
		return
	}

	resp = &dto.QueueRowDetail{
		ID:            data.ID,
		QueueNo:       data.QueueNo,
		MedicalRecord: data.MedicalRecord,
		LocationID:    data.LocationID,
		Type:          data.Type,
	}

	return
}

func (s *GlobalService) UpdateQueueById(ctx context.Context, id int, payload *dto.PayloadQueue) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	queue := &model.Queue{
		LocationID:    payload.LocationID,
		MedicalRecord: payload.MedicalRecord,
		QueueNo:       payload.QueueNo,
		Type:          payload.Type,
		ModifiedBy:    &user.Username,
	}

	if payload.MerchantID != nil {
		// ----- get merchant
		merchant, errr := s.globalRepository.FindMerchantById(ctx, int(*payload.MerchantID))
		if errr != nil {
			err = errr
			return
		}
		queue.MerchantName = merchant.Name
	}

	// ----- get location
	location, err := s.globalRepository.FindLocationById(ctx, int(payload.LocationID))
	if err != nil {
		return
	}
	queue.LocationName = location.Name
	resp, err = s.globalRepository.UpdateQueueById(ctx, id, queue)
	if err != nil {
		log.Errorf("err update Location %d", id)
		return
	}
	return
}

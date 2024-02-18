package impl

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/constant/role"
	statusadvertisement "github.com/maxzycon/rs-farmasi-be/pkg/constant/status_advertisement"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/timeutil"
	"gorm.io/datatypes"
)

func (s *GlobalService) GetAdvertisementPaginated(ctx context.Context, payload *dto.ParamsPaginationAdvertisement) (resp *dto.AdvertisementWrapper, err error) {
	// ----- Init response
	resp = &dto.AdvertisementWrapper{
		Summary: &dto.SummaryAdvertisement{
			OnGoing:    0,
			Finished:   0,
			OnSchedule: 0,
		},
		AdvertisementData: dto.AdvertisementData{
			Advertisements: make([]*dto.AdvertisementRow, 0),
			Paginator: dto.DefaultPaginationDtoRow{
				CurrentPage:   payload.Page,
				RecordPerPage: 0,
				LastPage:      0,
				TotalItem:     0,
			},
		},
	}

	user, _ := authutil.GetCredential(ctx)

	if payload.Limit < 1 {
		payload.Limit = 1
	}

	payload.Page = payload.Page - 1

	cond := squirrel.And{
		squirrel.Eq{
			"a.deleted_at": nil,
		},
	}

	if user.MerchantID != nil {
		cond = append(cond, squirrel.Eq{
			"a.merchant_id": user.MerchantID,
		})
	}

	if payload.Category > 0 {
		cond = append(cond, squirrel.Eq{
			"a.advertisement_category_id": payload.Category,
		})
	}

	if payload.Search != nil && *payload.Search != "" {
		cond = append(cond, squirrel.Eq{
			"a.name": payload.Search,
		})
	}

	// ----- Get data
	getStr, argStr, err := squirrel.Select(`a.id as id, 
			a.name, a.company, a.date_start, a.date_end, 
			a.merchant_id, m.name, a.advertisement_category_id, 
			ca.name,
			(CASE 
				WHEN (a.date_start >= CAST(NOW() as DATE) AND a.date_end <= CAST(NOW() as DATE))
					THEN "Berlangsung" 
				WHEN (a.date_start > CAST(NOW() as DATE) AND a.date_end > CAST(NOW() as DATE))
					THEN "Terjadwal" 
				WHEN (a.date_start <= CAST(NOW() as DATE) AND (a.date_end >= CAST(NOW() as DATE) OR a.date_end <= CAST(NOW() as DATE)))
					THEN "Selesai" 
				ELSE "-" 
			END) as status		  
		`).
		From("advertisements a").
		LeftJoin("advertisement_categories as ca ON ca.id = a.advertisement_category_id").
		LeftJoin("merchants as m ON m.id = a.merchant_id").
		Where(cond).
		Limit(payload.Limit).
		Offset(payload.Limit * payload.Page).
		OrderBy("a.id DESC").
		ToSql()

	if err != nil {
		return
	}

	row, err := s.db.WithContext(ctx).Raw(getStr, argStr...).Rows()
	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.AdvertisementRow{}
		var start time.Time
		var end time.Time
		err = row.Scan(
			&temp.ID, &temp.Name, &temp.Company,
			&start, &end, &temp.MerchantID,
			&temp.MerchantName, &temp.CategoryAdvertisementID,
			&temp.CategoryAdvertisementName, &temp.Status,
		)
		if err != nil {
			return
		}

		temp.DateStart = timeutil.ToStringDateOnly(start)
		temp.DateEnd = timeutil.ToStringDateOnly(end)
		resp.AdvertisementData.Advertisements = append(resp.AdvertisementData.Advertisements, &temp)
	}

	// ----- Count Pagination
	getStrCount, argStrCount, err := squirrel.Select(`COUNT(a.id) as id`).
		From("advertisements a").
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
		resp.AdvertisementData.Paginator.LastPage = uint64(total)
	} else {
		resp.AdvertisementData.Paginator.LastPage = uint64(total + 1)
	}
	resp.AdvertisementData.Paginator.RecordPerPage = payload.Limit
	resp.AdvertisementData.Paginator.TotalItem = uint64(totalRows)
	// ----- End count pagination

	// ----- Get Summary
	getSummaryStr, argSummaryStr, err := squirrel.Select(`COUNT(a.id), (CASE 
		WHEN (a.date_start >= CAST(NOW() as DATE) AND a.date_end <= CAST(NOW() as DATE))
				THEN "Berlangsung" 
		WHEN (a.date_start > CAST(NOW() as DATE) AND a.date_end > CAST(NOW() as DATE))
				THEN "Terjadwal" 
		WHEN (a.date_start <= CAST(NOW() as DATE) AND (a.date_end >= CAST(NOW() as DATE) OR a.date_end <= CAST(NOW() as DATE)))
				THEN "Selesai" 
		ELSE "-" 
		END) as status`).
		From("advertisements a").
		Where(cond).
		GroupBy(`(CASE 
			WHEN (a.date_start >= CAST(NOW() as DATE) AND a.date_end <= CAST(NOW() as DATE))
					THEN "Berlangsung" 
			WHEN (a.date_start > CAST(NOW() as DATE) AND a.date_end > CAST(NOW() as DATE))
					THEN "Terjadwal" 
			WHEN (a.date_start <= CAST(NOW() as DATE) AND (a.date_end >= CAST(NOW() as DATE) OR a.date_end <= CAST(NOW() as DATE)))
					THEN "Selesai" 
			ELSE "-" 
	END)`).
		ToSql()

	if err != nil {
		return
	}

	row, err = s.db.WithContext(ctx).Raw(getSummaryStr, argSummaryStr...).Rows()
	if err != nil {
		return
	}

	for row.Next() {
		var status string
		var count uint64
		err = row.Scan(
			&count, &status,
		)
		if err != nil {
			return
		}

		if status == statusadvertisement.FINISHED {
			resp.Summary.Finished = count
		}

		if status == statusadvertisement.ONGOING {
			resp.Summary.OnGoing = count
		}

		if status == statusadvertisement.ONSCHEDULE {
			resp.Summary.OnSchedule = count
		}
	}
	return
}

func (s *GlobalService) GetListContent(ctx context.Context, merchantIDstr string) (resp *dto.AdvertisementContentWrapper, err error) {
	resp = &dto.AdvertisementContentWrapper{
		Daily:    make([]*dto.AdvertisementListData, 0),
		Contents: make([]*dto.AdvertisementListData, 0),
	}

	// ----- Daily content
	cond := squirrel.And{
		squirrel.Eq{
			"a.deleted_at": nil,
		},
		squirrel.Eq{
			"m.id_str": merchantIDstr,
		},
		squirrel.GtOrEq{
			"CAST(NOW() AS DATE)": "a.date_start",
		},
		squirrel.LtOrEq{
			"CAST(NOW() AS DATE)": "a.date_start",
		},
	}

	dailyContent, args, err := squirrel.
		Select("a.id, a.document_path").
		From("advertisements as a").
		InnerJoin("merchants m ON m.id = a.merchant_id").
		Where(cond).
		ToSql()
	if err != nil {
		return
	}

	row, err := s.db.WithContext(ctx).Raw(dailyContent, args...).Rows()

	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.AdvertisementListData{}
		err = row.Scan(
			&temp.ID, &temp.URL,
		)

		if err != nil {
			return
		}
		temp.URL = s.conf.AWS_S3_URL + "/" + temp.URL
		resp.Daily = append(resp.Daily, &temp)
	}
	// ------ End daily

	// ----- All content
	content, args, err := squirrel.
		Select("a.id, a.document_path").
		From("advertisements as a").
		InnerJoin("merchants m ON m.id = a.merchant_id").
		Where(squirrel.And{
			squirrel.Eq{
				"a.deleted_at": nil,
			},
			squirrel.Eq{
				"m.id_str": merchantIDstr,
			},
		}).
		ToSql()
	if err != nil {
		return
	}

	row, err = s.db.WithContext(ctx).Raw(content, args...).Rows()

	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.AdvertisementListData{}
		err = row.Scan(
			&temp.ID, &temp.URL,
		)

		if err != nil {
			return
		}
		temp.URL = s.conf.AWS_S3_URL + "/" + temp.URL
		resp.Contents = append(resp.Daily, &temp)
	}
	// ----- End content

	return
}

func (s *GlobalService) GetAdvertisementPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllAdvertisement(ctx)
	if err != nil {
		log.Errorf("err get Advertisement paginated")
		return
	}
	resp = make([]*dto.DefaultPluck, 0)
	for _, row := range rows {
		resp = append(resp, &dto.DefaultPluck{
			ID:   row.ID,
			Name: row.Name,
		})
	}
	return
}

func (s *GlobalService) GetAdvertisementById(ctx context.Context, id int) (resp *dto.AdvertisementDetailRow, err error) {
	row, err := s.globalRepository.FindAdvertisementById(ctx, id)
	if err != nil {
		log.Errorf("err get Advertisement paginated")
		return
	}
	resp = &dto.AdvertisementDetailRow{
		ID:                        row.ID,
		Name:                      row.Name,
		Company:                   row.Company,
		MerchantID:                &row.Merchant.ID,
		MerchantName:              &row.Merchant.Name,
		CategoryAdvertisementID:   &row.AdvertisementCategoryID,
		CategoryAdvertisementName: &row.AdvertisementCategory.Name,
		Path:                      s.conf.AWS_S3_URL + "/" + row.DocumentPath,
	}

	d, err := row.DateEnd.Value()
	if err != nil {
		log.Errorf("err get Advertisement paginated")
		return
	}

	c, err := row.DateEnd.Value()
	if err != nil {
		log.Errorf("err get Advertisement paginated")
		return
	}
	r, _ := d.(time.Time)
	finalStart := timeutil.ToStringDateOnly(r)
	resp.DateStart = finalStart

	r, _ = c.(time.Time)
	end := timeutil.ToStringDateOnly(r)
	resp.DateEnd = end

	return
}

func (s *GlobalService) CreateAdvertisement(ctx context.Context, payload *dto.PayloadAdvertisement) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	entity := &model.Advertisement{
		Name:                    payload.Name,
		Company:                 payload.Company,
		MerchantID:              *payload.MerchantID,
		AdvertisementCategoryID: *payload.CategoryAdvertisementID,
		DocumentPath:            payload.DocumentPath,
		DateStart:               datatypes.Date(timeutil.ParseDate(payload.DateStart)),
		DateEnd:                 datatypes.Date(timeutil.ParseDate(payload.DateEnd)),
		Description:             payload.Description,
	}

	if user.Role != role.ROLE_OWNER {
		entity.MerchantID = *user.MerchantID
	}

	resp, err = s.globalRepository.CreateAdvertisement(ctx, entity)

	if err != nil {
		log.Errorf("err Advertisement status")
		return
	}
	return
}

func (s *GlobalService) UpdateAdvertisementById(ctx context.Context, id int, payload *dto.PayloadAdvertisement) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateAdvertisementById(ctx, id, &model.Advertisement{
		Name:                    payload.Name,
		Company:                 payload.Company,
		MerchantID:              *payload.MerchantID,
		AdvertisementCategoryID: *payload.CategoryAdvertisementID,
		DocumentPath:            payload.DocumentPath,
		DateStart:               datatypes.Date(timeutil.ParseDate(payload.DateStart)),
		DateEnd:                 datatypes.Date(timeutil.ParseDate(payload.DateEnd)),
		Description:             payload.Description,
	})
	if err != nil {
		log.Errorf("err update Advertisement %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteAdvertisementById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteAdvertisementById(ctx, id)
	if err != nil {
		log.Errorf("err delete Advertisement %d", id)
		return
	}
	return
}

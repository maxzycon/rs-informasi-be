package impl

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
)

func (s *GlobalService) GetMerchantPaginated(ctx context.Context, payload *dto.ParamsPaginationMerchant) (resp dto.MerchantWrapper, err error) {
	resp = dto.MerchantWrapper{
		SummaryMerchant: &dto.SummaryMerchant{
			TotalMerchant: 0,
		},
		Items: &dto.MerchantDataWrapper{
			Paginator: dto.DefaultPaginationDtoRow{
				CurrentPage:   payload.Page,
				RecordPerPage: payload.Limit,
				LastPage:      0,
				TotalItem:     0,
			},
			Merchants: make([]*dto.MerchantRowPaginated, 0),
		},
	}

	if payload.Limit < 1 {
		payload.Limit = 1
	}

	payload.Page = payload.Page - 1

	cond := squirrel.And{
		squirrel.Eq{
			"m.deleted_at": nil,
		},
	}

	if payload.Category > 0 {
		cond = append(cond, squirrel.Eq{
			"m.merchant_category_id": payload.Category,
		})
	}

	// ---- get data merchant
	getStr, argStr, err := squirrel.
		Select("m.id, m.name, m.email, m.pic_name, m.phone, m.photo, m.address, m.merchant_category_id, mc.name").
		From("merchants m").
		LeftJoin("merchant_categories mc ON mc.id = m.merchant_category_id").
		Where(cond).
		Limit(payload.Limit).
		Offset(payload.Limit * payload.Page).
		OrderBy(fmt.Sprintf("m.%s %s", payload.SortBy, payload.Order)).
		ToSql()

	if err != nil {
		return
	}

	row, err := s.db.WithContext(ctx).Raw(getStr, argStr...).Rows()
	if err != nil {
		return
	}

	for row.Next() {
		temp := dto.MerchantRowPaginated{}
		err = row.Scan(
			&temp.ID, &temp.Name, &temp.Email, &temp.PICName,
			&temp.Phone, &temp.Photo, &temp.Address, &temp.CategoryID,
			&temp.CategoryName,
		)

		if err != nil {
			return
		}

		if temp.Photo != nil {
			*temp.Photo = s.conf.AWS_S3_URL + "/" + *temp.Photo
		}

		resp.Items.Merchants = append(resp.Items.Merchants, &temp)
	}

	// ----- Count Pagination
	getStrCount, argStrCount, err := squirrel.Select(`COUNT(m.id) as id`).
		From("merchants m").
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

	resp.SummaryMerchant.TotalMerchant = uint64(totalRows)

	return
}

func (s *GlobalService) GetMerchantPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllMerchant(ctx)
	if err != nil {
		log.Errorf("err get Merchant paginated")
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

func (s *GlobalService) GetMerchantById(ctx context.Context, id int) (resp *dto.MerchantRow, err error) {
	row, err := s.globalRepository.FindMerchantById(ctx, id)
	if err != nil {
		log.Errorf("err get Merchant paginated")
		return
	}
	resp = &dto.MerchantRow{
		ID:                 row.ID,
		Name:               row.Name,
		Email:              row.Email,
		Phone:              row.Phone,
		PICName:            row.PICName,
		Photo:              row.Photo,
		Address:            row.Address,
		MerchantCategoryID: row.MerchantCategoryID,
	}

	if row.Photo != nil {
		*row.Photo = s.conf.AWS_S3_URL + "/" + *row.Photo
	}
	return
}

func (s *GlobalService) CreateMerchant(ctx context.Context, payload *dto.PayloadMerchant) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateMerchant(ctx, &model.Merchant{
		Name:               payload.Name,
		Address:            payload.Address,
		Phone:              payload.Phone,
		PICName:            payload.PICName,
		Email:              payload.Email,
		Photo:              payload.Photo,
		MerchantCategoryID: payload.MerchantCategoryID,
	})
	if err != nil {
		log.Errorf("err Merchant status")
		return
	}
	return
}

func (s *GlobalService) UpdateMerchantById(ctx context.Context, id int, payload *dto.PayloadMerchant) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateMerchantById(ctx, id, &model.Merchant{
		Name:               payload.Name,
		Address:            payload.Address,
		Phone:              payload.Phone,
		PICName:            payload.PICName,
		Email:              payload.Email,
		Photo:              payload.Photo,
		MerchantCategoryID: payload.MerchantCategoryID,
	})
	if err != nil {
		log.Errorf("err update Merchant %d", id)
		return
	}
	return
}

func (s *GlobalService) UpdateMerchantConfigById(ctx context.Context, id int, payload *dto.PayloadUpdateConfig) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateMerchantById(ctx, id, &model.Merchant{
		RunningText: payload.RunningText,
	})
	if err != nil {
		log.Errorf("err update Merchant %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteMerchantById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteMerchantById(ctx, id)
	if err != nil {
		log.Errorf("err delete Merchant %d", id)
		return
	}
	return
}

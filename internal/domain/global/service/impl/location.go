package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetLocationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindLocationPaginated(ctx, payload)
	if err != nil {
		log.Errorf("err get Location paginated")
		return
	}

	respToDto := make([]*dto.LocationRow, 0)
	list, ok := resp.Items.([]*model.Location)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.LocationRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetLocationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	user, _ := authutil.GetCredential(ctx)
	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("locations as l").
		Join("location_users as lu ON lu.location_id = l.id AND lu.user_id = ?", user.ID).
		ToSql()

	if err != nil {
		return
	}

	tx, err := s.db.Raw(parentSql, args...).Rows()

	if err != nil {
		return
	}

	for tx.Next() {
		tmp := dto.DefaultPluck{}
		err = tx.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			return
		}
		resp = append(resp, &dto.DefaultPluck{
			ID:   tmp.ID,
			Name: tmp.Name,
		})
	}

	return
}

func (s *GlobalService) GetLocationById(ctx context.Context, id int) (resp *dto.LocationRow, err error) {
	row, err := s.globalRepository.FindLocationById(ctx, id)
	if err != nil {
		log.Errorf("err get Location paginated")
		return
	}
	resp = &dto.LocationRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateLocation(ctx context.Context, payload *dto.PayloadLocation) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateLocation(ctx, &model.Location{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err Location status")
		return
	}
	return
}

func (s *GlobalService) UpdateLocationById(ctx context.Context, id int, payload *dto.PayloadLocation) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateLocationById(ctx, id, &model.Location{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err update Location %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteLocationById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteLocationById(ctx, id)
	if err != nil {
		log.Errorf("err delete Location %d", id)
		return
	}
	return
}

func (s *GlobalService) GetAllLocationByUser(ctx context.Context) (resp []*dto.LocationUserRow, err error) {
	resp = make([]*dto.LocationUserRow, 0)
	user, _ := authutil.GetCredential(ctx)
	parentSql, args, err := squirrel.
		Select("l.id, l.name, (CASE WHEN lu.id IS NOT NULL THEN 1 ELSE 0 END) as is_used").
		From("locations as l").
		LeftJoin("location_users as lu ON lu.location_id = l.id AND lu.user_id = ?", user.ID).
		ToSql()

	if err != nil {
		return
	}

	tx, err := s.db.Raw(parentSql, args...).Rows()

	if err != nil {
		return
	}

	for tx.Next() {
		tmp := dto.LocationUserRow{}
		err = tx.Scan(&tmp.ID, &tmp.Name, &tmp.IsUsed)
		if err != nil {
			return
		}
		resp = append(resp, &tmp)
	}
	return
}

func (s *GlobalService) UpdateAllLocationByUser(ctx context.Context, payload *dto.WrapperUpdateLocationUser) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	// ----- delete related user_id
	err = s.db.Unscoped().Where("user_id = ?", user.ID).Delete(&model.LocationUser{}).Error

	if err != nil {
		return nil, err
	}

	entities := make([]*model.LocationUser, 0)

	for _, v := range payload.Data {
		entities = append(entities, &model.LocationUser{
			UserID:     user.ID,
			LocationID: v.ID,
		})
	}

	err = s.db.Model(&model.LocationUser{}).Create(&entities).Error
	if err != nil {
		return nil, err
	}

	ok := int64(1)
	resp = &ok
	return
}

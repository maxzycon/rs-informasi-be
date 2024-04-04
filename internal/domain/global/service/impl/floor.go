package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetFloorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindFloorPaginated(ctx, payload)
	if err != nil {
		s.log.Errorf("err get Floor paginated")
		return
	}

	respToDto := make([]*dto.FloorRow, 0)
	list, ok := resp.Items.([]*model.Floor)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.FloorRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetFloorPluckByMerchantIdStr(ctx context.Context, merchantIdStr string) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("floors as l").
		LeftJoin("merchants as m ON m.id = l.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantIdStr,
			"l.deleted_at": nil,
		}).
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

func (s *GlobalService) GetFloorPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	user, _ := authutil.GetCredential(ctx)
	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("floors as l").
		Where(squirrel.Eq{
			"l.merchant_id": user.MerchantID,
			"l.deleted_at":  nil,
		}).
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

func (s *GlobalService) GetFloorById(ctx context.Context, id int) (resp *dto.FloorRow, err error) {
	row, err := s.globalRepository.FindFloorById(ctx, id)
	if err != nil {
		s.log.Errorf("err get Floor paginated")
		return
	}
	resp = &dto.FloorRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateFloor(ctx context.Context, payload *dto.PayloadFloor) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateFloor(ctx, &model.Floor{
		Name:       payload.Name,
		MerchantID: *user.MerchantID,
	})
	if err != nil {
		s.log.Errorf("err Floor status")
		return
	}
	return
}

func (s *GlobalService) UpdateFloorById(ctx context.Context, id int, payload *dto.PayloadFloor) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateFloorById(ctx, id, &model.Floor{
		Name: payload.Name,
	})
	if err != nil {
		s.log.Errorf("err update Floor %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteFloorById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteFloorById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Floor %d", id)
		return
	}
	return
}

func (s *GlobalService) GetAllFloorByUser(ctx context.Context) (resp []*dto.FloorUserRow, err error) {
	resp = make([]*dto.FloorUserRow, 0)
	user, _ := authutil.GetCredential(ctx)
	parentSql, args, err := squirrel.
		Select("l.id, l.name, (CASE WHEN lu.id IS NOT NULL THEN 1 ELSE 0 END) as is_used").
		From("Floors as l").
		LeftJoin("Floor_users as lu ON lu.Floor_id = l.id AND lu.user_id = ?", user.ID).
		ToSql()

	if err != nil {
		return
	}

	tx, err := s.db.Raw(parentSql, args...).Rows()

	if err != nil {
		return
	}

	for tx.Next() {
		tmp := dto.FloorUserRow{}
		err = tx.Scan(&tmp.ID, &tmp.Name, &tmp.IsUsed)
		if err != nil {
			return
		}
		resp = append(resp, &tmp)
	}
	return
}

package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetRoomPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindRoomPaginated(ctx, payload)
	if err != nil {
		s.log.Errorf("err get Room paginated")
		return
	}

	respToDto := make([]*dto.RoomRow, 0)
	list, ok := resp.Items.([]*model.Room)
	if ok {
		for _, v := range list {
			dto := &dto.RoomRow{
				ID:          v.ID,
				Name:        v.Name,
				FloorId:     v.FloorID,
				FloorName:   v.Floor.Name,
				Description: v.Description,
			}

			if v.Photo != nil {
				p := s.conf.AWS_S3_URL + "/" + *v.Photo
				dto.Photo = &p
			}

			respToDto = append(respToDto, dto)
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetRoomPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	user, _ := authutil.GetCredential(ctx)
	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("rooms as l").
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

func (s *GlobalService) GetRoomById(ctx context.Context, id int) (resp *dto.RoomRow, err error) {
	row, err := s.globalRepository.FindRoomById(ctx, id)
	if err != nil {
		s.log.Errorf("err get Room paginated")
		return
	}
	resp = &dto.RoomRow{
		ID:          row.ID,
		Name:        row.Name,
		FloorId:     row.Floor.ID,
		FloorName:   row.Floor.Name,
		Description: row.Description,
	}

	if row.Photo != nil {
		p := s.conf.AWS_S3_URL + "/" + *row.Photo
		resp.Photo = &p
	}

	return
}

func (s *GlobalService) CreateRoom(ctx context.Context, payload *dto.PayloadRoom) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateRoom(ctx, &model.Room{
		Name:        payload.Name,
		MerchantID:  *user.MerchantID,
		FloorID:     payload.FloorId,
		Description: payload.Description,
		Photo:       payload.Photo,
	})
	if err != nil {
		s.log.Errorf("err Room status")
		return
	}
	return
}

func (s *GlobalService) UpdateRoomById(ctx context.Context, id int, payload *dto.PayloadRoom) (resp *int64, err error) {
	row, err := s.GetRoomById(ctx, id)
	if err != nil {
		s.log.Errorf("err update Information %d", id)
		return
	}

	entity := &model.Room{
		Name:        payload.Name,
		FloorID:     payload.FloorId,
		Description: payload.Description,
	}

	if payload.Photo != nil && row.Photo != payload.Photo {
		entity.Photo = payload.Photo
	}

	resp, err = s.globalRepository.UpdateRoomById(ctx, id, entity)

	if err != nil {
		s.log.Errorf("err update Room %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteRoomById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteRoomById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Room %d", id)
		return
	}
	return
}

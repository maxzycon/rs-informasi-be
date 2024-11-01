package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindRoomById(ctx context.Context, id int) (resp *model.Room, err error) {
	tx := r.db.WithContext(ctx).Preload("Floor").First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllRoom(ctx context.Context) (resp []*model.Room, err error) {
	resp = make([]*model.Room, 0)
	tx := r.db.WithContext(ctx).Model(&model.Room{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindRoomPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Rooms []*model.Room = make([]*model.Room, 0)
	user, err := authutil.GetCredential(ctx)

	if err != nil {
		return
	}

	sql := r.db.Debug().WithContext(ctx).Preload("Floor")
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	sql.Scopes(payload.Pagination(&Rooms, &resp.Paginator, sql)).Find(&Rooms)
	resp.Items = Rooms
	return
}

func (r *GlobalRepository) CreateRoom(ctx context.Context, entity *model.Room) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Room{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateRoomById(ctx context.Context, id int, entity *model.Room) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteRoomById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Room{}, id)
	return &tx.RowsAffected, tx.Error
}

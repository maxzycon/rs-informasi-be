package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindFloorById(ctx context.Context, id int) (resp *model.Floor, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllFloor(ctx context.Context) (resp []*model.Floor, err error) {
	resp = make([]*model.Floor, 0)
	tx := r.db.WithContext(ctx).Model(&model.Floor{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindFloorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Floors []*model.Floor = make([]*model.Floor, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Floors, &resp.Paginator, sql)).Find(&Floors)
	resp.Items = Floors
	return
}

func (r *GlobalRepository) CreateFloor(ctx context.Context, entity *model.Floor) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Floor{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateFloorById(ctx context.Context, id int, entity *model.Floor) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteFloorById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Floor{}, id)
	return &tx.RowsAffected, tx.Error
}

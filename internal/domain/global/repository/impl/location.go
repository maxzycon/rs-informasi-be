package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindLocationById(ctx context.Context, id int) (resp *model.Location, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllLocation(ctx context.Context) (resp []*model.Location, err error) {
	resp = make([]*model.Location, 0)
	tx := r.db.WithContext(ctx).Model(&model.Location{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindLocationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Locations []*model.Location = make([]*model.Location, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Locations, &resp.Paginator, sql)).Find(&Locations)
	resp.Items = Locations
	return
}

func (r *GlobalRepository) CreateLocation(ctx context.Context, entity *model.Location) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Location{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateLocationById(ctx context.Context, id int, entity *model.Location) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteLocationById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Location{}, id)
	return &tx.RowsAffected, tx.Error
}

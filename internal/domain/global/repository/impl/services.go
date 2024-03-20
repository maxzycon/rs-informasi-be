package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindServicesById(ctx context.Context, id int) (resp *model.Services, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllServices(ctx context.Context) (resp []*model.Services, err error) {
	resp = make([]*model.Services, 0)
	tx := r.db.WithContext(ctx).Model(&model.Services{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindServicesPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Servicess []*model.Services = make([]*model.Services, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Servicess, &resp.Paginator, sql)).Find(&Servicess)
	resp.Items = Servicess
	return
}

func (r *GlobalRepository) CreateServices(ctx context.Context, entity *model.Services) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Services{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateServicesById(ctx context.Context, id int, entity *model.Services) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteServicesById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Services{}, id)
	return &tx.RowsAffected, tx.Error
}

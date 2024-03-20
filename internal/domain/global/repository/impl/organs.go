package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindOrganById(ctx context.Context, id int) (resp *model.Organ, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllOrgan(ctx context.Context) (resp []*model.Organ, err error) {
	resp = make([]*model.Organ, 0)
	tx := r.db.WithContext(ctx).Model(&model.Organ{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindOrganPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Organs []*model.Organ = make([]*model.Organ, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Organs, &resp.Paginator, sql)).Find(&Organs)
	resp.Items = Organs
	return
}

func (r *GlobalRepository) CreateOrgan(ctx context.Context, entity *model.Organ) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Organ{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateOrganById(ctx context.Context, id int, entity *model.Organ) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteOrganById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Organ{}, id)
	return &tx.RowsAffected, tx.Error
}

package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindAdvertisementCategoryById(ctx context.Context, id int) (resp *model.AdvertisementCategory, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllAdvertisementCategory(ctx context.Context) (resp []*model.AdvertisementCategory, err error) {
	resp = make([]*model.AdvertisementCategory, 0)
	tx := r.db.WithContext(ctx).Model(&model.AdvertisementCategory{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAdvertisementCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var AdvertisementCategorys []*model.AdvertisementCategory = make([]*model.AdvertisementCategory, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&AdvertisementCategorys, &resp.Paginator, sql)).Find(&AdvertisementCategorys)
	resp.Items = AdvertisementCategorys
	return
}

func (r *GlobalRepository) CreateAdvertisementCategory(ctx context.Context, entity *model.AdvertisementCategory) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.AdvertisementCategory{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateAdvertisementCategoryById(ctx context.Context, id int, entity *model.AdvertisementCategory) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteAdvertisementCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.AdvertisementCategory{}, id)
	return &tx.RowsAffected, tx.Error
}

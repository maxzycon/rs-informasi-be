package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindMerchantCategoryById(ctx context.Context, id int) (resp *model.MerchantCategory, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllMerchantCategory(ctx context.Context) (resp []*model.MerchantCategory, err error) {
	resp = make([]*model.MerchantCategory, 0)
	tx := r.db.WithContext(ctx).Model(&model.MerchantCategory{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var MerchantCategorys []*model.MerchantCategory = make([]*model.MerchantCategory, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&MerchantCategorys, &resp.Paginator, sql)).Find(&MerchantCategorys)
	resp.Items = MerchantCategorys
	return
}

func (r *GlobalRepository) CreateMerchantCategory(ctx context.Context, entity *model.MerchantCategory) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.MerchantCategory{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateMerchantCategoryById(ctx context.Context, id int, entity *model.MerchantCategory) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.MerchantCategory{}, id)
	return &tx.RowsAffected, tx.Error
}

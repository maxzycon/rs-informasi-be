package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindProductCategoryById(ctx context.Context, id int) (resp *model.ProductCategory, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllProductCategory(ctx context.Context) (resp []*model.ProductCategory, err error) {
	user, err := authutil.GetCredential(ctx)
	if err != nil {
		return
	}

	resp = make([]*model.ProductCategory, 0)
	sql := r.db.WithContext(ctx).Model(&model.ProductCategory{})

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	tx := sql.Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindProductCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var ProductCategorys []*model.ProductCategory = make([]*model.ProductCategory, 0)
	user, err := authutil.GetCredential(ctx)

	if err != nil {
		return
	}

	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}
	sql.Scopes(payload.Pagination(&ProductCategorys, &resp.Paginator, sql)).Find(&ProductCategorys)
	resp.Items = ProductCategorys
	return
}

func (r *GlobalRepository) CreateProductCategory(ctx context.Context, entity *model.ProductCategory) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.ProductCategory{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateProductCategoryById(ctx context.Context, id int, entity *model.ProductCategory) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteProductCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.ProductCategory{}, id)
	return &tx.RowsAffected, tx.Error
}

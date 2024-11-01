package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindProductById(ctx context.Context, id int) (resp *model.Product, err error) {
	tx := r.db.WithContext(ctx).Preload("Detail").Preload("ProductCategory").First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllProduct(ctx context.Context) (resp []*model.Product, err error) {
	resp = make([]*model.Product, 0)
	tx := r.db.WithContext(ctx).Model(&model.Product{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindProductPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Products []*model.Product = make([]*model.Product, 0)
	sql := r.db.Debug().WithContext(ctx).Preload("Detail").Preload("ProductCategory")
	user, err := authutil.GetCredential(ctx)

	if err != nil {
		return
	}
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}
	sql.Scopes(payload.Pagination(&Products, &resp.Paginator, sql)).Find(&Products)
	resp.Items = Products
	return
}

func (r *GlobalRepository) CreateProduct(ctx context.Context, entity *model.Product) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Product{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateProductById(ctx context.Context, id int, entity *model.Product) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteProductById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Product{}, id)
	return &tx.RowsAffected, tx.Error
}

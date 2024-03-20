package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindMerchantById(ctx context.Context, id int) (resp *model.Merchant, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllMerchant(ctx context.Context) (resp []*model.Merchant, err error) {
	resp = make([]*model.Merchant, 0)
	tx := r.db.WithContext(ctx).Model(&model.Merchant{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindMerchantPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Merchants []*model.Merchant = make([]*model.Merchant, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Merchants, &resp.Paginator, sql)).Find(&Merchants)
	resp.Items = Merchants
	return
}

func (r *GlobalRepository) CreateMerchant(ctx context.Context, entity *model.Merchant) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Merchant{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateMerchantById(ctx context.Context, id int, entity *model.Merchant) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteMerchantById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Merchant{}, id)
	return &tx.RowsAffected, tx.Error
}

package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindAdvertisementById(ctx context.Context, id int) (resp *model.Advertisement, err error) {
	tx := r.db.WithContext(ctx).Preload("Merchant").Preload("AdvertisementCategory").First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllAdvertisement(ctx context.Context) (resp []*model.Advertisement, err error) {
	user, err := authutil.GetCredential(ctx)
	if err != nil {
		return
	}

	resp = make([]*model.Advertisement, 0)
	sql := r.db.WithContext(ctx).Model(&model.Advertisement{})

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	tx := sql.Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAdvertisementPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Advertisements []*model.Advertisement = make([]*model.Advertisement, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Advertisements, &resp.Paginator, sql)).Find(&Advertisements)
	resp.Items = Advertisements
	return
}

func (r *GlobalRepository) CreateAdvertisement(ctx context.Context, entity *model.Advertisement) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Advertisement{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateAdvertisementById(ctx context.Context, id int, entity *model.Advertisement) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteAdvertisementById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Advertisement{}, id)
	return &tx.RowsAffected, tx.Error
}

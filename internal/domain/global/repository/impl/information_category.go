package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindInformationCategoryById(ctx context.Context, id int) (resp *model.InformationCategory, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllInformationCategory(ctx context.Context) (resp []*model.InformationCategory, err error) {
	user, err := authutil.GetCredential(ctx)
	if err != nil {
		return
	}

	resp = make([]*model.InformationCategory, 0)
	sql := r.db.WithContext(ctx).Model(&model.InformationCategory{})

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	tx := sql.Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindInformationCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var InformationCategorys []*model.InformationCategory = make([]*model.InformationCategory, 0)
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
	sql.Scopes(payload.Pagination(&InformationCategorys, &resp.Paginator, sql)).Find(&InformationCategorys)
	resp.Items = InformationCategorys
	return
}

func (r *GlobalRepository) CreateInformationCategory(ctx context.Context, entity *model.InformationCategory) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.InformationCategory{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateInformationCategoryById(ctx context.Context, id int, entity *model.InformationCategory) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteInformationCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.InformationCategory{}, id)
	return &tx.RowsAffected, tx.Error
}

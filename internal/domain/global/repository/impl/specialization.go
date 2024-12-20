package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindMerchantSpecializationById(ctx context.Context, id int) (resp *model.Specialization, err error) {
	tx := r.db.WithContext(ctx).Preload("Organ").First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllMerchantSpecialization(ctx context.Context) (resp []*model.Specialization, err error) {
	user, err := authutil.GetCredential(ctx)
	if err != nil {
		return
	}

	resp = make([]*model.Specialization, 0)

	sql := r.db.WithContext(ctx).Model(&model.Specialization{})

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	tx := sql.Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindMerchantSpecializationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var MerchantSpecializations []*model.Specialization = make([]*model.Specialization, 0)
	user, err := authutil.GetCredential(ctx)

	if err != nil {
		return
	}
	sql := r.db.Debug().WithContext(ctx).Preload("Organ")
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	sql.Scopes(payload.Pagination(&MerchantSpecializations, &resp.Paginator, sql)).Find(&MerchantSpecializations)
	resp.Items = MerchantSpecializations
	return
}

func (r *GlobalRepository) CreateMerchantSpecialization(ctx context.Context, entity *model.Specialization) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Specialization{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateMerchantSpecializationById(ctx context.Context, id int, entity *model.Specialization) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteMerchantSpecializationById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Specialization{}, id)
	return &tx.RowsAffected, tx.Error
}

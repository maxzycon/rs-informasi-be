package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindFacilityById(ctx context.Context, id int) (resp *model.Facility, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllFacility(ctx context.Context) (resp []*model.Facility, err error) {
	user, err := authutil.GetCredential(ctx)
	if err != nil {
		return
	}

	resp = make([]*model.Facility, 0)
	sql := r.db.WithContext(ctx).Model(&model.Facility{})

	if user.Role == uint(role.ROLE_ADMIN) {
		sql = sql.Where("merchant_id = ?", user.MerchantID)
	}

	tx := sql.Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindFacilityPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Facilitys []*model.Facility = make([]*model.Facility, 0)
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

	sql.Scopes(payload.Pagination(&Facilitys, &resp.Paginator, sql)).Find(&Facilitys)
	resp.Items = Facilitys
	return
}

func (r *GlobalRepository) CreateFacility(ctx context.Context, entity *model.Facility) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Facility{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateFacilityById(ctx context.Context, id int, entity *model.Facility) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteFacilityById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Facility{}, id)
	return &tx.RowsAffected, tx.Error
}

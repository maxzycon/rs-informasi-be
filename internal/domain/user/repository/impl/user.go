package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
	"gorm.io/gorm"
)

func (r *UserRepository) FindAllUserPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload, claims *authutil.UserClaims) (resp pagination.DefaultPagination, err error) {
	var users []*model.User = make([]*model.User, 0)
	sql := r.db.WithContext(ctx)

	if payload.Order == "" {
		payload.Order = "desc"
	}

	if payload.SortBy == "" {
		payload.SortBy = "id"
	}

	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql.Where("username LIKE ?", search)
	}
	sql.Scopes(payload.PaginationV2(&resp.Paginator)).Find(&users)
	resp.Items = users
	return
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {
	resp = &model.User{}
	tx := r.db.Where("id_agen = ?", username).First(&resp)
	return resp, tx.Error
}

func (r *UserRepository) FindByIdAndDepartmentId(ctx context.Context, id int, departmentId uint) (resp *model.User, err error) {
	resp = &model.User{}
	tx := r.db.Where("id = ?", id).First(&resp)
	return resp, tx.Error
}

func (r UserRepository) FindUserByUsernameLogin(ctx context.Context, username string) (resp *model.User, err error) {
	resp = &model.User{}
	tx := r.db.Where("username = ?", username).Preload("Merchant").First(&resp)
	return resp, tx.Error
}

func (r *UserRepository) Create(ctx context.Context, payload *model.User) (resp *model.User, err error) {
	tx := r.db.Create(&payload)
	return payload, tx.Error
}

func (r *UserRepository) Update(ctx context.Context, payload *model.User, id int) (resp *model.User, err error) {
	payload.ID = uint(id)
	tx := r.db.Model(&payload).Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx)

	tx.Updates(&payload)
	return payload, tx.Error
}

func (r *UserRepository) FindById(ctx context.Context, id int) (resp *model.User, err error) {
	resp = &model.User{}
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *UserRepository) DeleteUserById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.Delete(&model.User{}, id)
	resp = &tx.RowsAffected
	return resp, tx.Error
}

func (r *UserRepository) GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error) {
	tx := r.db.WithContext(ctx).First(&resp, userId)
	return resp, tx.Error
}

func (r *UserRepository) UpdatePasswordByUserId(ctx context.Context, id int, newPassword *string) (resp *int64, err error) {
	tx := r.db.Model(&model.User{}).WithContext(ctx).Where("id", id).Update("password", newPassword)
	resp = &tx.RowsAffected
	return resp, tx.Error
}

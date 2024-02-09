package repository

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"

	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

type UserRepository interface {
	FindById(ctx context.Context, id int) (resp *model.User, err error)
	FindByIdAndDepartmentId(ctx context.Context, id int, departmentId uint) (resp *model.User, err error)
	FindUserByUsernameLogin(ctx context.Context, username string) (resp *model.User, err error)
	FindUserByUsername(ctx context.Context, username string) (resp *model.User, err error)
	FindAllUserPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload, claims *authutil.UserClaims) (resp pagination.DefaultPagination, err error)
	Create(ctx context.Context, payload *model.User) (resp *model.User, err error)
	Update(ctx context.Context, payload *model.User, id int) (resp *model.User, err error)
	DeleteUserById(ctx context.Context, id int) (resp *int64, err error)
	UpdatePasswordByUserId(ctx context.Context, id int, newPassword *string) (resp *int64, err error)
	GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error)
}

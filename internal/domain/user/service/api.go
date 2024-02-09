package service

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/internal/domain/user/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"

	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

type UserService interface {
	// ---- Users
	GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error)
	UpdateUserProfile(ctx context.Context, id int, password string) (resp *int64, err error)
	GetById(ctx context.Context, id int) (resp *dto.UserRowDetail, err error)
	DeleteUserById(ctx context.Context, id int, claims *authutil.UserClaims) (resp *int64, err error)
	Login(ctx context.Context, payload dto.PayloadLogin) (resp *dto.LoginRes, err error)
	GetUserByUsername(ctx context.Context, username string) (resp *model.User, err error)
	UpdateUser(ctx context.Context, payload dto.PayloadUpdateUser, id int, claims *authutil.UserClaims) (resp *model.User, err error)
	CreateUser(ctx context.Context, payload dto.PayloadCreateUser, claims *authutil.UserClaims) (resp *model.User, err error)
	GetUserPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload, claims *authutil.UserClaims) (resp pagination.DefaultPagination, err error)
}

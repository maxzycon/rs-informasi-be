package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/maxzycon/rs-informasi-be/internal/domain/user/service"
	"github.com/maxzycon/rs-informasi-be/pkg/constant/role"
	"github.com/maxzycon/rs-informasi-be/pkg/middleware"
)

const (
	Create = "/users"

	Update            = "/users/:id"
	GetById           = "/users/:id"
	Delete            = "/users/:id"
	GetUsersPaginated = "/users"
	UpdateUserProfile = "/profile"

	Login        = "/login"
	GetUserLogin = "/user"
)

type UsersControllerParams struct {
	V1          fiber.Router
	Conf        *config.Config
	UserService service.UserService
	Middleware  middleware.GlobalMiddleware
}
type usersController struct {
	v1          fiber.Router
	conf        *config.Config
	userService service.UserService
	middleware  middleware.GlobalMiddleware
}

func New(params *UsersControllerParams) *usersController {
	return &usersController{
		v1:          params.V1,
		conf:        params.Conf,
		userService: params.UserService,
		middleware:  params.Middleware,
	}
}
func (pc *usersController) Init() {
	pc.v1.Post(Create, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerCreateUser)
	pc.v1.Put(Update, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerUpdateUser)
	pc.v1.Get(GetUserLogin, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN, role.ROLE_FARMASI, role.ROLE_KASIR, role.ROLE_MARKETING, role.ROLE_MULTIMEDIA}), pc.handlerUser)
	pc.v1.Get(GetUsersPaginated, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetUsersPaginated)
	pc.v1.Delete(Delete, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerDeleteUserById)
	pc.v1.Get(GetById, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handlerGetUserById)

	pc.v1.Put(UpdateUserProfile, pc.middleware.Protected([]uint{role.ROLE_OWNER, role.ROLE_SUPER_ADMIN}), pc.handerUpdateUserProfile)

	pc.v1.Post(Login, pc.handlerLogin)
}

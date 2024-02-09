package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/user/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/constant"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (c *usersController) handlerUser(f *fiber.Ctx) (err error) {
	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.WriteSuccessResponse(f, user)
}

func (c *usersController) handerUpdateUserProfile(f *fiber.Ctx) (err error) {
	payload := dto.PayloadUpdateProfile{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleUpdateUserProfile] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.UpdateUserProfile(f.Context(), int(user.ID), payload.Password)
	if err != nil {
		log.Error("err update user profile controller")
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *usersController) handlerCreateUser(f *fiber.Ctx) (err error) {
	payload := dto.PayloadCreateUser{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleCreateUser] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.userService.CreateUser(f.Context(), payload, user)

	if err != nil {
		log.Errorf("[user.go][hadnlerCreateUser] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, constant.Success)
}

func (c *usersController) handlerUpdateUser(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handlerGetUserById] err parse params")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateUser{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleCreateUser] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.userService.UpdateUser(f.Context(), payload, id, user)

	if err != nil {
		log.Errorf("[user.go][hadnlerCreateUser] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, constant.Success)
}

func (c *usersController) handlerGetUsersPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handlerGetUsersPaginated] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}
	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}

	resp, err := c.userService.GetUserPaginated(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("[user.go][handlerGetUsersPaginated] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *usersController) handlerGetUserById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handlerGetUserById] err parse params")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.GetById(f.Context(), id)

	if err != nil {
		log.Errorf("[user.go][handlerGetUserById] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *usersController) handlerDeleteUserById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handlerDeleteUserById] err parse params")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}

	resp, err := c.userService.DeleteUserById(f.Context(), id, user)

	if err != nil {
		log.Errorf("[user.go][handlerDeleteUserById] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

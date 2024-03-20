package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-informasi-be/internal/domain/user/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
)

func (c *usersController) handlerLogin(f *fiber.Ctx) (err error) {
	payload := dto.PayloadLogin{}
	err = f.BodyParser(&payload)
	if err != nil {
		log.Errorf("[handlerLogin] err body parse")
		err = errors.ErrBadRequest
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.Login(f.Context(), payload)

	if err != nil {
		log.Errorf("[handlerLogin] err service at controller %v", err)
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.WriteSuccessResponse(f, resp)
}

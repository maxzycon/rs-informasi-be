package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetAllUserPluck(f *fiber.Ctx) (err error) {
	user, err := authutil.GetCredential(f.Context())
	if err != nil {
		c.log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}

	resp, err := c.globalService.GetAllUserPluck(f.Context(), user)
	if err != nil {
		c.log.Errorf("[user.go][handlerGetAllUserPluck] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

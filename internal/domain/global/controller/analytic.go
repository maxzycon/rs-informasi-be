package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetDashboardAnalytic(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetDashboardAnalytic(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Location pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

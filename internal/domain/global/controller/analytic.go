package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetDashboardAnalytic(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetDashboardAnalytic(f.Context())
	if err != nil {
		log.Errorf("err service at controller Location pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

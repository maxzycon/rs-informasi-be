package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllServicePluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetServicePluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Service pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateService(f *fiber.Ctx) (err error) {
	payload := dto.PayloadService{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Service")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateService(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Service :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateService(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Service")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadService{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Service")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateServiceById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Service :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetServicePaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Service")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetServicePaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Service :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetServiceById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Service get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetServiceById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Service get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteService(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Service delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteServiceById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Service delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

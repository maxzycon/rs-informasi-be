package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllOrganPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetOrganPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Organ pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateOrgan(f *fiber.Ctx) (err error) {
	payload := dto.PayloadOrgan{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Organ")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateOrgan(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Organ :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateOrgan(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Organ")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadOrgan{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Organ")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateOrganById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Organ :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetOrganPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Organ")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetOrganPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Organ :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetOrganById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Organ get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetOrganById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Organ get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteOrgan(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Organ delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteOrganById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Organ delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

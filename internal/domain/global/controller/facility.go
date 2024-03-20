package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllFacilityPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetFacilityPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Facility pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateFacility(f *fiber.Ctx) (err error) {
	payload := dto.PayloadFacility{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Facility")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateFacility(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Facility :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateFacility(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Facility")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadFacility{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Facility")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateFacilityById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Facility :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetFacilityPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Facility")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetFacilityPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Facility :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetFacilityById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Facility get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetFacilityById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Facility get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteFacility(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Facility delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteFacilityById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Facility delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

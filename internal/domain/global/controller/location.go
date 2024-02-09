package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllLocationPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetLocationPluck(f.Context())
	if err != nil {
		log.Errorf("err service at controller Location pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateLocation(f *fiber.Ctx) (err error) {
	payload := dto.PayloadLocation{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Location")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateLocation(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create Location :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateLocation(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Location")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadLocation{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Location")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateLocationById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Location :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetLocationPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body paginated Location")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetLocationPaginated(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller paginated Location :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetLocationById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Location get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetLocationById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Location get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteLocation(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Location delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteLocationById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Location delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

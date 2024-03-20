package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllFloorPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetFloorPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Floor pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllFloorUser(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetAllFloorByUser(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Floor user:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateFloor(f *fiber.Ctx) (err error) {
	payload := dto.PayloadFloor{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Floor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateFloor(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Floor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateFloor(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Floor")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadFloor{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Floor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateFloorById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Floor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetFloorPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Floor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetFloorPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Floor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetFloorById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Floor get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetFloorById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Floor get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteFloor(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Floor delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteFloorById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Floor delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

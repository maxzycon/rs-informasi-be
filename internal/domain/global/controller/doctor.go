package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllDoctorPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetDoctorPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Doctor pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateDoctor(f *fiber.Ctx) (err error) {
	payload := dto.PayloadDoctor{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Doctor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateDoctor(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Doctor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateDoctor(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Doctor")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadDoctor{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Doctor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateDoctorById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Doctor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetDoctorPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Doctor")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetDoctorPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Doctor :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetDoctorById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Doctor get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetDoctorById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Doctor get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteDoctor(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Doctor delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteDoctorById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Doctor delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

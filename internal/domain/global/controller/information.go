package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllInformationPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetInformationPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Information pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateInformation(f *fiber.Ctx) (err error) {
	payload := dto.PayloadInformation{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Information")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateInformation(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Information :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateInformation(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Information")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadInformation{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Information")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateInformationById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Information :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetInformationPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Information")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetInformationPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Information :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetInformationById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Information get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetInformationById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Information get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteInformation(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Information delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteInformationById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Information delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllInformationCategoryPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetInformationCategoryPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller InformationCategory pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateInformationCategory(f *fiber.Ctx) (err error) {
	payload := dto.PayloadInformationCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create InformationCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateInformationCategory(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create InformationCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateInformationCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update InformationCategory")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadInformationCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update InformationCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateInformationCategoryById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update InformationCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetInformationCategoryPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated InformationCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetInformationCategoryPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated InformationCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetInformationCategoryById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params InformationCategory get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetInformationCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller InformationCategory get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteInformationCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params InformationCategory delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteInformationCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller InformationCategory delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

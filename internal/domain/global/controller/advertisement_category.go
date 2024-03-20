package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllAdvertisementCategoryPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetAdvertisementCategoryPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller AdvertisementCategory pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateAdvertisementCategory(f *fiber.Ctx) (err error) {
	payload := dto.PayloadAdvertisementCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create AdvertisementCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateAdvertisementCategory(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create AdvertisementCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateAdvertisementCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update AdvertisementCategory")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadAdvertisementCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update AdvertisementCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateAdvertisementCategoryById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update AdvertisementCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetAdvertisementCategoryPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated AdvertisementCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetAdvertisementCategoryPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated AdvertisementCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAdvertisementCategoryById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params AdvertisementCategory get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetAdvertisementCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller AdvertisementCategory get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteAdvertisementCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params AdvertisementCategory delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteAdvertisementCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller AdvertisementCategory delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

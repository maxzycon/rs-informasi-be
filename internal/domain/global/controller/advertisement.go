package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetAllAdvertisementPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetAdvertisementPluck(f.Context())
	if err != nil {
		log.Errorf("err service at controller Advertisement pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateAdvertisement(f *fiber.Ctx) (err error) {
	payload := dto.PayloadAdvertisement{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Advertisement")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateAdvertisement(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create Advertisement :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateAdvertisement(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Advertisement")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadAdvertisement{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Advertisement")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateAdvertisementById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Advertisement :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetAdvertisementPaginated(f *fiber.Ctx) (err error) {
	payload := dto.ParamsPaginationAdvertisement{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body paginated Advertisement")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetAdvertisementPaginated(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller paginated Advertisement :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAdvertisementById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Advertisement get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetAdvertisementById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Advertisement get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteAdvertisement(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Advertisement delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteAdvertisementById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Advertisement delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerContentAdvertisement(f *fiber.Ctx) (err error) {
	id := f.Params("id")
	if id == "" {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Advertisement delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetListContent(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Advertisement delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllMerchantSpecializationPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetMerchantSpecializationPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller MerchantSpecialization pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateMerchantSpecialization(f *fiber.Ctx) (err error) {
	payload := dto.PayloadMerchantSpecialization{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create MerchantSpecialization")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateMerchantSpecialization(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create MerchantSpecialization :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateMerchantSpecialization(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update MerchantSpecialization")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadMerchantSpecialization{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update MerchantSpecialization")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateMerchantSpecializationById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update MerchantSpecialization :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetMerchantSpecializationPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated MerchantSpecialization")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantSpecializationPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated MerchantSpecialization :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMerchantSpecializationById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params MerchantSpecialization get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantSpecializationById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller MerchantSpecialization get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteMerchantSpecialization(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params MerchantSpecialization delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteMerchantSpecializationById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller MerchantSpecialization delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

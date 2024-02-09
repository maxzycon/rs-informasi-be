package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllMerchantCategoryPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetMerchantCategoryPluck(f.Context())
	if err != nil {
		log.Errorf("err service at controller MerchantCategory pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateMerchantCategory(f *fiber.Ctx) (err error) {
	payload := dto.PayloadMerchantCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create MerchantCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateMerchantCategory(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create MerchantCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateMerchantCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update MerchantCategory")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadMerchantCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update MerchantCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateMerchantCategoryById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update MerchantCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetMerchantCategoryPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body paginated MerchantCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantCategoryPaginated(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller paginated MerchantCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMerchantCategoryById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params MerchantCategory get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantCategoryById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller MerchantCategory get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteMerchantCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params MerchantCategory delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteMerchantCategoryById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller MerchantCategory delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

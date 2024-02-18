package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetAllMerchantPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetMerchantPluck(f.Context())
	if err != nil {
		log.Errorf("err service at controller Merchant pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateMerchant(f *fiber.Ctx) (err error) {
	payload := dto.PayloadMerchant{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Merchant")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateMerchant(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create Merchant :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateMerchant(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Merchant")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadMerchant{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Merchant")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateMerchantById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Merchant :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateMerchantConfig(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f.Context())

	payload := dto.PayloadUpdateConfig{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update config merchant")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateMerchantConfigById(f.Context(), int(*user.MerchantID), &payload)

	if err != nil {
		log.Errorf("err service at controller update config merchant :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetMerchantPaginated(f *fiber.Ctx) (err error) {
	payload := dto.ParamsPaginationMerchant{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body paginated Merchant")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantPaginated(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller paginated Merchant :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMerchantById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Merchant get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetMerchantById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Merchant get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetRunningTextByMerchantIdStr(f *fiber.Ctx) (err error) {
	id := f.Params("id")
	if id == "" {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Merchant get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetRunningTextByMerchantIdStr(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Merchant get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteMerchant(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Merchant delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteMerchantById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Merchant delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

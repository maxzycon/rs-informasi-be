package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
)

func (c *GlobalController) handlerGetDashboard(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}
	resp, err := c.globalService.GetDashboardKiosk(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information list kiosk:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllInformationKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}
	categoryId := f.QueryInt("category_id", 0)
	resp, err := c.globalService.GetInformationListKiosk(f.Context(), categoryId, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information list kiosk:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetInformationKioskById(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	id := f.Params("id")
	if id == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetInformationKiosk(f.Context(), id, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information kiosk by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllFacilitiesKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetFacilitiesListKiosk(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetFacilityKioskById(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	id := f.Params("id")
	if id == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetFacilitieskioskById(f.Context(), id, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information kiosk by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllRoomsKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	floorId := f.QueryInt("floor_id", 0)

	resp, err := c.globalService.GetRoomsListKiosk(f.Context(), floorId, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetRoomKioskById(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	id := f.Params("id")
	if id == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetRoomskioskById(f.Context(), id, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information kiosk by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllServicesKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetServicesListKiosk(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetServiceKioskById(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	id := f.Params("id")
	if id == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetServicekioskById(f.Context(), id, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller information kiosk by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllProductsKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	categoryId := f.QueryInt("category_id", 0)

	resp, err := c.globalService.GetProductsListKiosk(f.Context(), categoryId, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllDoctorsKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	specializationID := f.QueryInt("specialization_id", 0)
	day := f.QueryInt("day", 0)

	resp, err := c.globalService.GetDoctorsListKiosk(f.Context(), specializationID, day, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetDoctorKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}
	id := f.Params("id")

	resp, err := c.globalService.GetDoctorsKiosk(f.Context(), id, merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMasterFloorKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetFloorPluckByMerchantIdStr(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMasterSpecializationKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetMerchantSpecializationByMerchantStrIdPluck(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMasterCategoryProductKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetProductCategoryPluckByMerchantStrId(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetMasterCategoryInformationKiosk(f *fiber.Ctx) (err error) {
	merchantUUID := f.Query("merchant_id")
	if merchantUUID == "" {
		return httputil.WriteErrorResponse(f, errors.ErrBadRequest)
	}

	resp, err := c.globalService.GetInformationCategoryPluckByMerchantStrId(f.Context(), merchantUUID)
	if err != nil {
		c.log.Errorf("err service at controller facilities kiosk pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

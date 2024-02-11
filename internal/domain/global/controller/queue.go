package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
)

func (c *GlobalController) handlerDisplayQueue(f *fiber.Ctx) (err error) {
	payload := dto.ParamsQueueDisplay{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetDashboardDisplay(f.Context(), &payload, payload.MerchantIdStr)

	if err != nil {
		log.Errorf("err service at controller create Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerQueueBySearch(f *fiber.Ctx) (err error) {
	search := f.Query("search")
	if search == "" {
		err = errors.ErrBadRequest
		log.Errorf("err parse queries search")
		return httputil.WriteErrorResponse(f, err)
	}

	merchantId := f.Query("merchant")
	if merchantId == "" {
		err = errors.ErrBadRequest
		log.Errorf("err parse queries search")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetQueueBySearch(f.Context(), merchantId, search)

	if err != nil {
		log.Errorf("err service at controller create Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerUpdateFollowUpPhone(f *fiber.Ctx) (err error) {
	id := f.Params("id")
	if id == "" {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Queue")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateFollowUpQueue{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	err = c.globalService.UpdateFuQueueNo(f.Context(), id, payload.NewFollowUpPhoneNo)

	if err != nil {
		log.Errorf("err service at controller update Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	var resp int64 = 1
	return httputil.WriteSuccessResponseAffectedRow(f, &resp)
}

func (c *GlobalController) handlerCreateQueue(f *fiber.Ctx) (err error) {
	payload := dto.PayloadQueue{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateQueue(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateStatusQueue(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Queue")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateQueue{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateStatusQueueById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateQueueById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Queue")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadQueue{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateQueueById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetQueuePaginated(f *fiber.Ctx) (err error) {
	payload := dto.ParamsQueueQueries{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body paginated Queue")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetQueuePaginated(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller paginated Queue :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

// func (c *GlobalController) handlerGetQueueById(f *fiber.Ctx) (err error) {
// 	id, err := f.ParamsInt("id")
// 	if err != nil {
// 		err = errors.ErrBadRequest
// 		log.Errorf("err parse params Queue get by id")
// 		return httputil.WriteErrorResponse(f, err)
// 	}
// 	resp, err := c.globalService.GetQueueById(f.Context(), id)

// 	if err != nil {
// 		log.Errorf("err service at controller Queue get by id:%+v", err)
// 		return httputil.WriteErrorResponse(f, err)
// 	}

// 	return httputil.WriteSuccessResponse(f, resp)
// }

func (c *GlobalController) handlerDeleteQueue(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Queue delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteQueueById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Queue delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllRoomPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetRoomPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Room pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateRoom(f *fiber.Ctx) (err error) {
	payload := dto.PayloadRoom{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Room")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateRoom(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Room :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateRoom(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Room")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadRoom{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Room")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateRoomById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Room :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetRoomPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Room")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetRoomPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Room :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetRoomById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Room get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetRoomById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Room get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteRoom(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Room delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteRoomById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Room delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

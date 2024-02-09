package httputil

import (
	"net/http"

	"github.com/maxzycon/rs-farmasi-be/pkg/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-farmasi-be/pkg/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
)

func WriteSuccessResponse(e *fiber.Ctx, payload interface{}) error {
	return WriteResponse(e, dto.ResponseParam{
		Status: http.StatusOK,
		Payload: dto.BaseResponse{
			Data: payload,
		},
	})
}

func WriteSuccessResponseAffectedRow(e *fiber.Ctx, affectedRow *int64) error {
	if *affectedRow > 0 {
		return WriteResponse(e, dto.ResponseParam{
			Status: http.StatusOK,
			Payload: dto.BaseResponse{
				Data: constant.Success,
			},
		})
	}
	return WriteErrorResponse(e, errors.ErrInternalServer)
}

func WriteErrorResponse(e *fiber.Ctx, er error) error {
	errResp := errors.GetErrorResponse(er)
	return WriteResponse(e, dto.ResponseParam{
		Status: int(errResp.HTTPCode),
		Payload: dto.BaseResponse{
			Error: &dto.ErrorResponse{
				Code:    errResp.Code,
				Message: errResp.Message,
			},
		},
	})
}

func WriteResponse(e *fiber.Ctx, param dto.ResponseParam) error {
	return e.Status(param.Status).JSON(param.Payload)
}

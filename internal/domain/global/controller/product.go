package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllProductPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetProductPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller Product pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateProduct(f *fiber.Ctx) (err error) {
	payload := dto.PayloadProduct{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create Product")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateProduct(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateProduct(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update Product")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadProduct{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update Product")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateProductById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetProductPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated Product")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetProductPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetProductById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Product get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetProductById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Product get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteProduct(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params Product delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteProductById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller Product delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

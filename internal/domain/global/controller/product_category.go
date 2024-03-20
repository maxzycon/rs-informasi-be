package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/errors"
	"github.com/maxzycon/rs-informasi-be/pkg/httputil"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (c *GlobalController) handlerGetAllProductCategoryPluck(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetProductCategoryPluck(f.Context())
	if err != nil {
		c.log.Errorf("err service at controller ProductCategory pluck :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerCreateProductCategory(f *fiber.Ctx) (err error) {
	payload := dto.PayloadProductCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body create ProductCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateProductCategory(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller create ProductCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateProductCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params update ProductCategory")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadProductCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body update ProductCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateProductCategoryById(f.Context(), id, &payload)

	if err != nil {
		c.log.Errorf("err service at controller update ProductCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetProductCategoryPaginated(f *fiber.Ctx) (err error) {
	payload := pagination.DefaultPaginationPayload{}
	err = f.QueryParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse body paginated ProductCategory")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetProductCategoryPaginated(f.Context(), &payload)

	if err != nil {
		c.log.Errorf("err service at controller paginated ProductCategory :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetProductCategoryById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params ProductCategory get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetProductCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller ProductCategory get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteProductCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		c.log.Errorf("err parse params ProductCategory delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteProductCategoryById(f.Context(), id)

	if err != nil {
		c.log.Errorf("err service at controller ProductCategory delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

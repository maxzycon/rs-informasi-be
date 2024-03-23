package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
	"github.com/maxzycon/rs-informasi-be/pkg/util/timeutil"
	"gorm.io/datatypes"
)

func (s *GlobalService) GetProductPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindProductPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get Product paginated")
		return
	}

	respToDto := make([]*dto.ProductRow, 0)
	list, ok := resp.Items.([]*model.Product)
	if ok {
		for _, v := range list {

			e := &dto.ProductRow{
				ID:                v.ID,
				Name:              v.Name,
				CategoryProductID: v.ProductCategoryID,
				CategoryName:      v.ProductCategory.Name,
				Price:             v.Price,
				IsDiscount:        v.IsDiscount,
				AmountDiscount:    v.AmountDiscount,
				DetailProduct:     make([]*dto.DetailProductRow, 0),
			}

			if v.Photo != nil {
				s := s.conf.AWS_S3_URL + "/" + *v.Photo
				e.Photo = &s
			}

			if v.IsDiscount {
				e.StartDiscount = v.DiscountStartDate
				e.EndDiscount = v.DiscountEndDate
				// ----- add amount, start,end date
			}

			for _, r := range v.Detail {
				e.DetailProduct = append(e.DetailProduct, &dto.DetailProductRow{
					ID:          r.ID,
					Description: r.Name,
				})
			}

			respToDto = append(respToDto, e)
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetProductPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllProduct(ctx)
	if err != nil {
		s.log.Error("err get Product paginated")
		return
	}
	resp = make([]*dto.DefaultPluck, 0)
	for _, row := range rows {
		resp = append(resp, &dto.DefaultPluck{
			ID:   row.ID,
			Name: row.Name,
		})
	}
	return
}

func (s *GlobalService) GetProductById(ctx context.Context, id int) (resp *dto.ProductRow, err error) {
	row, err := s.globalRepository.FindProductById(ctx, id)
	if err != nil {
		s.log.Error("err get Product paginated")
		return
	}
	resp = &dto.ProductRow{
		ID:                row.ID,
		Name:              row.Name,
		CategoryProductID: row.ProductCategoryID,
		CategoryName:      row.ProductCategory.Name,
		Price:             row.Price,
		IsDiscount:        row.IsDiscount,
		AmountDiscount:    row.AmountDiscount,
		DetailProduct:     make([]*dto.DetailProductRow, 0),
	}

	if row.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *row.Photo
		resp.Photo = &s
	}

	if row.IsDiscount {
		resp.StartDiscount = row.DiscountStartDate
		resp.EndDiscount = row.DiscountEndDate
		// ----- add amount, start,end date
	}

	for _, v := range row.Detail {
		resp.DetailProduct = append(resp.DetailProduct, &dto.DetailProductRow{
			ID:          v.ID,
			Description: v.Name,
		})
	}

	return
}

func (s *GlobalService) CreateProduct(ctx context.Context, payload *dto.PayloadProduct) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	entity := &model.Product{
		Name:              payload.Name,
		IsDiscount:        payload.IsDiscount,
		Price:             payload.Price,
		ProductCategoryID: payload.CategoryProductID,
		AmountDiscount:    payload.AmountDiscount,
		Photo:             payload.Photo,
		MerchantID:        *user.MerchantID,
	}

	if payload.IsDiscount && payload.StartDiscount != nil && payload.EndDiscount != nil {
		start, _ := timeutil.FromString(*payload.StartDiscount)
		end, _ := timeutil.FromString(*payload.EndDiscount)
		s := datatypes.Date(start)
		e := datatypes.Date(end)
		entity.DiscountStartDate = &s
		entity.DiscountEndDate = &e
	}

	for _, v := range payload.DetailProduct {
		entity.Detail = append(entity.Detail, model.DetailProduct{
			Name: v.Description,
		})
	}

	resp, err = s.globalRepository.CreateProduct(ctx, entity)
	if err != nil {
		s.log.Error("err Product status")
		return
	}
	return
}

func (s *GlobalService) UpdateProductById(ctx context.Context, id int, payload *dto.PayloadProduct) (resp *int64, err error) {
	row, err := s.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	entity := &model.Product{
		Name:              payload.Name,
		IsDiscount:        payload.IsDiscount,
		Price:             payload.Price,
		ProductCategoryID: payload.CategoryProductID,
		AmountDiscount:    payload.AmountDiscount,
	}

	if payload.Photo != nil && row.Photo != payload.Photo {
		entity.Photo = payload.Photo
	}

	if payload.IsDiscount && payload.StartDiscount != nil && payload.EndDiscount != nil {
		start, _ := timeutil.FromString(*payload.StartDiscount)
		end, _ := timeutil.FromString(*payload.EndDiscount)
		s := datatypes.Date(start)
		e := datatypes.Date(end)
		entity.DiscountStartDate = &s
		entity.DiscountEndDate = &e
	}

	err = s.db.Model(&model.DetailProduct{}).Unscoped().Where("product_id = ?", id).Delete(&model.DetailProduct{}).Error
	if err != nil {
		return nil, err
	}

	for _, v := range payload.DetailProduct {
		entity.Detail = append(entity.Detail, model.DetailProduct{
			Name: v.Description,
		})
	}

	resp, err = s.globalRepository.UpdateProductById(ctx, id, entity)
	if err != nil {
		s.log.Errorf("err update Product %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteProductById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteProductById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Product %d", id)
		return
	}
	return
}

package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetProductCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindProductCategoryPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get ProductCategory paginated")
		return
	}

	respToDto := make([]*dto.ProductCategoryRow, 0)
	list, ok := resp.Items.([]*model.ProductCategory)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.ProductCategoryRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetProductCategoryPluckByMerchantStrId(ctx context.Context, merchantStrID string) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("product_categories as l").
		LeftJoin("merchants as m ON m.id = l.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantStrID,
			"l.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		return
	}

	tx, err := s.db.Raw(parentSql, args...).Rows()

	if err != nil {
		return
	}

	for tx.Next() {
		tmp := dto.DefaultPluck{}
		err = tx.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			return
		}
		resp = append(resp, &dto.DefaultPluck{
			ID:   tmp.ID,
			Name: tmp.Name,
		})
	}

	return
}

func (s *GlobalService) GetProductCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllProductCategory(ctx)
	if err != nil {
		s.log.Error("err get ProductCategory paginated")
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

func (s *GlobalService) GetProductCategoryById(ctx context.Context, id int) (resp *dto.ProductCategoryRow, err error) {
	row, err := s.globalRepository.FindProductCategoryById(ctx, id)
	if err != nil {
		s.log.Error("err get ProductCategory paginated")
		return
	}
	resp = &dto.ProductCategoryRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateProductCategory(ctx context.Context, payload *dto.PayloadProductCategory) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateProductCategory(ctx, &model.ProductCategory{
		Name:       payload.Name,
		MerchantID: *user.MerchantID,
	})
	if err != nil {
		s.log.Error("err ProductCategory status")
		return
	}
	return
}

func (s *GlobalService) UpdateProductCategoryById(ctx context.Context, id int, payload *dto.PayloadProductCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateProductCategoryById(ctx, id, &model.ProductCategory{
		Name: payload.Name,
	})
	if err != nil {
		s.log.Errorf("err update ProductCategory %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteProductCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteProductCategoryById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete ProductCategory %d", id)
		return
	}
	return
}

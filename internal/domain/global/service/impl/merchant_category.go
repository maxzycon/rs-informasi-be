package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetMerchantCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindMerchantCategoryPaginated(ctx, payload)
	if err != nil {
		log.Errorf("err get MerchantCategory paginated")
		return
	}

	respToDto := make([]*dto.MerchantCategoryRow, 0)
	list, ok := resp.Items.([]*model.MerchantCategory)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.MerchantCategoryRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetMerchantCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllMerchantCategory(ctx)
	if err != nil {
		log.Errorf("err get MerchantCategory paginated")
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

func (s *GlobalService) GetMerchantCategoryById(ctx context.Context, id int) (resp *dto.MerchantCategoryRow, err error) {
	row, err := s.globalRepository.FindMerchantCategoryById(ctx, id)
	if err != nil {
		log.Errorf("err get MerchantCategory paginated")
		return
	}
	resp = &dto.MerchantCategoryRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateMerchantCategory(ctx context.Context, payload *dto.PayloadMerchantCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateMerchantCategory(ctx, &model.MerchantCategory{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err MerchantCategory status")
		return
	}
	return
}

func (s *GlobalService) UpdateMerchantCategoryById(ctx context.Context, id int, payload *dto.PayloadMerchantCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateMerchantCategoryById(ctx, id, &model.MerchantCategory{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err update MerchantCategory %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteMerchantCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteMerchantCategoryById(ctx, id)
	if err != nil {
		log.Errorf("err delete MerchantCategory %d", id)
		return
	}
	return
}

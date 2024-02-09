package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetAdvertisementCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindAdvertisementCategoryPaginated(ctx, payload)
	if err != nil {
		log.Errorf("err get AdvertisementCategory paginated")
		return
	}

	respToDto := make([]*dto.AdvertisementCategoryRow, 0)
	list, ok := resp.Items.([]*model.AdvertisementCategory)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.AdvertisementCategoryRow{
				ID:          v.ID,
				Name:        v.Name,
				Description: v.Description,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetAdvertisementCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllAdvertisementCategory(ctx)
	if err != nil {
		log.Errorf("err get AdvertisementCategory paginated")
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

func (s *GlobalService) GetAdvertisementCategoryById(ctx context.Context, id int) (resp *dto.AdvertisementCategoryRow, err error) {
	row, err := s.globalRepository.FindAdvertisementCategoryById(ctx, id)
	if err != nil {
		log.Errorf("err get AdvertisementCategory paginated")
		return
	}
	resp = &dto.AdvertisementCategoryRow{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description,
	}
	return
}

func (s *GlobalService) CreateAdvertisementCategory(ctx context.Context, payload *dto.PayloadAdvertisementCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateAdvertisementCategory(ctx, &model.AdvertisementCategory{
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		log.Errorf("err AdvertisementCategory status")
		return
	}
	return
}

func (s *GlobalService) UpdateAdvertisementCategoryById(ctx context.Context, id int, payload *dto.PayloadAdvertisementCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateAdvertisementCategoryById(ctx, id, &model.AdvertisementCategory{
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		log.Errorf("err update AdvertisementCategory %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteAdvertisementCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteAdvertisementCategoryById(ctx, id)
	if err != nil {
		log.Errorf("err delete AdvertisementCategory %d", id)
		return
	}
	return
}

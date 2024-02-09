package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetLocationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindLocationPaginated(ctx, payload)
	if err != nil {
		log.Errorf("err get Location paginated")
		return
	}

	respToDto := make([]*dto.LocationRow, 0)
	list, ok := resp.Items.([]*model.Location)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.LocationRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetLocationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllLocation(ctx)
	if err != nil {
		log.Errorf("err get Location paginated")
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

func (s *GlobalService) GetLocationById(ctx context.Context, id int) (resp *dto.LocationRow, err error) {
	row, err := s.globalRepository.FindLocationById(ctx, id)
	if err != nil {
		log.Errorf("err get Location paginated")
		return
	}
	resp = &dto.LocationRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateLocation(ctx context.Context, payload *dto.PayloadLocation) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateLocation(ctx, &model.Location{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err Location status")
		return
	}
	return
}

func (s *GlobalService) UpdateLocationById(ctx context.Context, id int, payload *dto.PayloadLocation) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateLocationById(ctx, id, &model.Location{
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("err update Location %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteLocationById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteLocationById(ctx, id)
	if err != nil {
		log.Errorf("err delete Location %d", id)
		return
	}
	return
}

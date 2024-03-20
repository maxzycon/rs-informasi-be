package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetOrganPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindOrganPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get Organ paginated")
		return
	}

	respToDto := make([]*dto.OrganRow, 0)
	list, ok := resp.Items.([]*model.Organ)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.OrganRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetOrganPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllOrgan(ctx)
	if err != nil {
		s.log.Error("err get Organ paginated")
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

func (s *GlobalService) GetOrganById(ctx context.Context, id int) (resp *dto.OrganRow, err error) {
	row, err := s.globalRepository.FindOrganById(ctx, id)
	if err != nil {
		s.log.Error("err get Organ paginated")
		return
	}
	resp = &dto.OrganRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateOrgan(ctx context.Context, payload *dto.PayloadOrgan) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateOrgan(ctx, &model.Organ{
		Name: payload.Name,
	})
	if err != nil {
		s.log.Error("err Organ status")
		return
	}
	return
}

func (s *GlobalService) UpdateOrganById(ctx context.Context, id int, payload *dto.PayloadOrgan) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateOrganById(ctx, id, &model.Organ{
		Name: payload.Name,
	})
	if err != nil {
		s.log.Errorf("err update Organ %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteOrganById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteOrganById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Organ %d", id)
		return
	}
	return
}

package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetFacilityPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindFacilityPaginated(ctx, payload)
	if err != nil {
		s.log.Errorf("err get Facility paginated")
		return
	}

	respToDto := make([]*dto.FacilityRow, 0)
	list, ok := resp.Items.([]*model.Facility)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.FacilityRow{
				ID:          v.ID,
				Name:        v.Name,
				Description: v.Desc,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetFacilityPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllFacility(ctx)
	if err != nil {
		s.log.Errorf("err get Facility paginated")
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

func (s *GlobalService) GetFacilityById(ctx context.Context, id int) (resp *dto.FacilityRow, err error) {
	row, err := s.globalRepository.FindFacilityById(ctx, id)
	if err != nil {
		s.log.Errorf("err get Facility paginated")
		return
	}
	resp = &dto.FacilityRow{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Desc,
	}
	return
}

func (s *GlobalService) CreateFacility(ctx context.Context, payload *dto.PayloadFacility) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateFacility(ctx, &model.Facility{
		Name:       payload.Name,
		Desc:       &payload.Description,
		MerchantID: *user.MerchantID,
	})
	if err != nil {
		s.log.Errorf("err Facility status")
		return
	}
	return
}

func (s *GlobalService) UpdateFacilityById(ctx context.Context, id int, payload *dto.PayloadFacility) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateFacilityById(ctx, id, &model.Facility{
		Name: payload.Name,
		Desc: &payload.Description,
	})
	if err != nil {
		s.log.Errorf("err update Facility %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteFacilityById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteFacilityById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Facility %d", id)
		return
	}
	return
}

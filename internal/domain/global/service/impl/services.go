package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetServicePaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindServicesPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get Services paginated")
		return
	}

	respToDto := make([]*dto.ServiceRow, 0)
	list, ok := resp.Items.([]*model.Services)
	if ok {
		for _, v := range list {
			dto := &dto.ServiceRow{
				ID:   v.ID,
				Name: v.Name,
				Desc: v.Desc,
			}

			if v.Photo != nil {
				p := s.conf.AWS_S3_URL + "/" + *v.Photo
				dto.Photo = &p
			}

			respToDto = append(respToDto, dto)
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetServicePluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllServices(ctx)
	if err != nil {
		s.log.Error("err get Services paginated")
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

func (s *GlobalService) GetServiceById(ctx context.Context, id int) (resp *dto.ServiceRow, err error) {
	row, err := s.globalRepository.FindServicesById(ctx, id)
	if err != nil {
		s.log.Error("err get Services paginated")
		return
	}
	resp = &dto.ServiceRow{
		ID:   row.ID,
		Name: row.Name,
		Desc: row.Desc,
	}

	if row.Photo != nil {
		p := s.conf.AWS_S3_URL + "/" + *row.Photo
		resp.Photo = &p
	}

	return
}

func (s *GlobalService) CreateService(ctx context.Context, payload *dto.PayloadService) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateServices(ctx, &model.Services{
		Name:       payload.Name,
		Desc:       payload.Desc,
		Photo:      payload.Photo,
		MerchantID: *user.MerchantID,
	})
	if err != nil {
		s.log.Error("err Services status")
		return
	}
	return
}

func (s *GlobalService) UpdateServiceById(ctx context.Context, id int, payload *dto.PayloadService) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateServicesById(ctx, id, &model.Services{
		Name:  payload.Name,
		Desc:  payload.Desc,
		Photo: payload.Photo,
	})
	if err != nil {
		s.log.Errorf("err update Services %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteServiceById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteServicesById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Services %d", id)
		return
	}
	return
}

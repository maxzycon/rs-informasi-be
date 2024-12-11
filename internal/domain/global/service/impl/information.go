package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetInformationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindInformationPaginated(ctx, payload)
	if err != nil {
		s.log.Errorf("err get Information paginated")
		return
	}

	respToDto := make([]*dto.InformationRow, 0)
	list, ok := resp.Items.([]*model.Information)
	if ok {
		for _, v := range list {
			dto := &dto.InformationRow{
				ID:                    v.ID,
				Name:                  v.Name,
				InformationCategoryID: v.InformationCategoryID,
				Desc:                  v.Desc,
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

func (s *GlobalService) GetInformationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllInformation(ctx)
	if err != nil {
		s.log.Errorf("err get Information paginated")
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

func (s *GlobalService) GetInformationById(ctx context.Context, id int) (resp *dto.InformationRow, err error) {
	row, err := s.globalRepository.FindInformationById(ctx, id)
	if err != nil {
		s.log.Errorf("err get Information paginated")
		return
	}
	resp = &dto.InformationRow{
		ID:                    row.ID,
		Name:                  row.Name,
		InformationCategoryID: row.InformationCategoryID,
		Desc:                  row.Desc,
	}

	if row.Photo != nil {
		p := s.conf.AWS_S3_URL + "/" + *row.Photo
		resp.Photo = &p
	}
	return
}

func (s *GlobalService) CreateInformation(ctx context.Context, payload *dto.PayloadInformation) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateInformation(ctx, &model.Information{
		Name:                  payload.Name,
		Desc:                  payload.Desc,
		Photo:                 payload.Photo,
		InformationCategoryID: payload.InformationCategoryID,
		MerchantID:            *user.MerchantID,
	})
	if err != nil {
		s.log.Errorf("err Information status")
		return
	}
	return
}

func (s *GlobalService) UpdateInformationById(ctx context.Context, id int, payload *dto.PayloadInformation) (resp *int64, err error) {
	row, err := s.GetInformationById(ctx, id)
	if err != nil {
		s.log.Errorf("err update Information %d", id)
		return
	}

	entity := &model.Information{
		Name:                  payload.Name,
		Desc:                  payload.Desc,
		InformationCategoryID: payload.InformationCategoryID,
	}

	if payload.Photo != nil && *row.Photo != *payload.Photo {
		entity.Photo = payload.Photo
	}

	resp, err = s.globalRepository.UpdateInformationById(ctx, id, entity)
	if err != nil {
		s.log.Errorf("err update Information %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteInformationById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteInformationById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Information %d", id)
		return
	}
	return
}

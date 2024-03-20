package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetInformationCategoryPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindInformationCategoryPaginated(ctx, payload)
	if err != nil {
		s.log.Errorf("err get InformationCategory paginated")
		return
	}

	respToDto := make([]*dto.InformationCategoryRow, 0)
	list, ok := resp.Items.([]*model.InformationCategory)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.InformationCategoryRow{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetInformationCategoryPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllInformationCategory(ctx)
	if err != nil {
		s.log.Errorf("err get InformationCategory paginated")
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

func (s *GlobalService) GetInformationCategoryById(ctx context.Context, id int) (resp *dto.InformationCategoryRow, err error) {
	row, err := s.globalRepository.FindInformationCategoryById(ctx, id)
	if err != nil {
		s.log.Errorf("err get InformationCategory paginated")
		return
	}
	resp = &dto.InformationCategoryRow{
		ID:   row.ID,
		Name: row.Name,
	}
	return
}

func (s *GlobalService) CreateInformationCategory(ctx context.Context, payload *dto.PayloadInformationCategory) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateInformationCategory(ctx, &model.InformationCategory{
		Name:       payload.Name,
		MerchantID: *user.MerchantID,
	})
	if err != nil {
		s.log.Errorf("err InformationCategory status")
		return
	}
	return
}

func (s *GlobalService) UpdateInformationCategoryById(ctx context.Context, id int, payload *dto.PayloadInformationCategory) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateInformationCategoryById(ctx, id, &model.InformationCategory{
		Name: payload.Name,
	})
	if err != nil {
		s.log.Errorf("err update InformationCategory %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteInformationCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteInformationCategoryById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete InformationCategory %d", id)
		return
	}
	return
}

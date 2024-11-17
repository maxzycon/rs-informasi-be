package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetMerchantSpecializationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindMerchantSpecializationPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get MerchantSpecialization paginated")
		return
	}

	respToDto := make([]*dto.MerchantSpecializationRow, 0)
	list, ok := resp.Items.([]*model.Specialization)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.MerchantSpecializationRow{
				ID:   v.ID,
				Name: v.Name,
				// OrganID:   v.OrganID,
				// OrganName: v.Organ.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetMerchantSpecializationByMerchantStrIdPluck(ctx context.Context, merchantStrID string) (resp []*dto.DefaultPluck, err error) {
	resp = make([]*dto.DefaultPluck, 0)

	parentSql, args, err := squirrel.
		Select("l.id, l.name").
		From("specializations as l").
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

func (s *GlobalService) GetMerchantSpecializationPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllMerchantSpecialization(ctx)
	if err != nil {
		s.log.Error("err get MerchantSpecialization paginated")
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

func (s *GlobalService) GetMerchantSpecializationById(ctx context.Context, id int) (resp *dto.MerchantSpecializationRow, err error) {
	row, err := s.globalRepository.FindMerchantSpecializationById(ctx, id)
	if err != nil {
		s.log.Error("err get MerchantSpecialization paginated")
		return
	}
	resp = &dto.MerchantSpecializationRow{
		ID:   row.ID,
		Name: row.Name,
		// OrganID:   row.OrganID,
		// OrganName: row.Organ.Name,
	}
	return
}

func (s *GlobalService) CreateMerchantSpecialization(ctx context.Context, payload *dto.PayloadMerchantSpecialization) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	resp, err = s.globalRepository.CreateMerchantSpecialization(ctx, &model.Specialization{
		Name:       payload.Name,
		MerchantID: *user.MerchantID,
		// OrganID:    payload.OrganID,
	})
	if err != nil {
		s.log.Error("err MerchantSpecialization status")
		return
	}
	return
}

func (s *GlobalService) UpdateMerchantSpecializationById(ctx context.Context, id int, payload *dto.PayloadMerchantSpecialization) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateMerchantSpecializationById(ctx, id, &model.Specialization{
		Name: payload.Name,
		// OrganID: payload.OrganID,
	})
	if err != nil {
		s.log.Errorf("err update MerchantSpecialization %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteMerchantSpecializationById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteMerchantSpecializationById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete MerchantSpecialization %d", id)
		return
	}
	return
}

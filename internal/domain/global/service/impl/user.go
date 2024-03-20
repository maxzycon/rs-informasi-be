package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
)

func (s *GlobalService) GetAllUserPluck(ctx context.Context, claims *authutil.UserClaims) (resp []*dto.UserRowPluck, err error) {
	rows, err := s.globalRepository.FindAllUser(ctx, claims)
	if err != nil {
		s.log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}

	resp = make([]*dto.UserRowPluck, 0)
	for _, v := range rows {
		resp = append(resp, &dto.UserRowPluck{
			ID:   int(v.ID),
			Name: v.Name,
			Role: v.Role,
		})
	}

	return
}

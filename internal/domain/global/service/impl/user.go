package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
)

func (s *GlobalService) GetAllUserPluck(ctx context.Context, claims *authutil.UserClaims) (resp []*dto.UserRowPluck, err error) {
	rows, err := s.globalRepository.FindAllUser(ctx, claims)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
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

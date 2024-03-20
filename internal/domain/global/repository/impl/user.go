package impl

import (
	"context"

	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
)

func (r *GlobalRepository) FindAllUser(ctx context.Context, claims *authutil.UserClaims) (resp []*model.User, err error) {
	resp = make([]*model.User, 0)
	query := r.db
	tx := query.Find(&resp)
	return resp, tx.Error
}

package authutil

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	ID            uint       `json:"id"`
	Username      string     `json:"username"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Phone         *string    `json:"phone"`
	Photo         *string    `json:"photo"`
	MerchantID    *uint      `json:"merchant_id"`
	MerchantIDStr *uuid.UUID `json:"merchant_id_str"`
	Role          uint       `json:"role"`
}

func GetCredential(f context.Context) (resp *UserClaims, err error) {
	resp, ok := f.Value("user").(*UserClaims)
	if !ok {
		log.Errorf("err parse data profile to userClaims")
		err = errors.ErrUnauthorized
		return
	}
	return
}

package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxzycon/rs-farmasi-be/internal/config"
	UserService "github.com/maxzycon/rs-farmasi-be/internal/domain/user/service/impl"
	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
	"golang.org/x/exp/slices"
)

type GlobalMiddleware struct {
	UserService *UserService.UserService
	Conf        *config.Config
}

func (m *GlobalMiddleware) Protected(roleAccess []uint) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(m.Conf.JWT_SECRET_KEY)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return httputil.WriteErrorResponse(c, errors.ErrUnauthorized)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// ---- JWT
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			id, ok := claims["id"].(float64)
			if !ok {
				return httputil.WriteErrorResponse(ctx, errors.ErrUnauthorized)
			}

			// --- get user by id
			userRow, err := m.UserService.GetUserByIdToken(ctx.Context(), uint(id))
			if err != nil {
				log.Errorf("err getUserByIdToken %d", uint(id))
				return httputil.WriteErrorResponse(ctx, errors.ErrUnauthorized)
			}

			if !slices.Contains(roleAccess, userRow.Role) {
				return httputil.WriteErrorResponse(ctx, errors.ErrForbidden)
			}

			resp := &authutil.UserClaims{
				ID:         userRow.ID,
				Name:       userRow.Name,
				Phone:      &userRow.Phone,
				Username:   userRow.Username,
				Email:      userRow.Email,
				Photo:      userRow.Photo,
				Role:       userRow.Role,
				MerchantID: userRow.MerchantID,
			}
			// ----- set user to all request with protected
			ctx.Context().SetUserValue("user", resp)
			// ctx.Locals("user_profile", resp)
			return ctx.Next()
		},
	})
}

package impl

import (
	"context"
	"strings"
	"time"

	"github.com/maxzycon/rs-farmasi-be/pkg/authutil"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/user/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/errors"
	"github.com/maxzycon/rs-farmasi-be/pkg/helper"
	"github.com/maxzycon/rs-farmasi-be/pkg/model"
	"github.com/maxzycon/rs-farmasi-be/pkg/util/pagination"
)

func (s *UserService) Login(ctx context.Context, payload dto.PayloadLogin) (resp *dto.LoginRes, err error) {
	user, err := s.UserRepository.FindUserByUsernameLogin(ctx, payload.Username)
	if err != nil {
		log.Errorf("[Login] findUserByusername :%+v", err)
		return
	}

	password := helper.CheckPasswordHash(payload.Password, user.Password)
	if !password {
		log.Errorf("[Login] err hash password doens't match")
		err = errors.ErrInvalidPassword
		return
	}

	// --- set 30 day exp
	exp := time.Now().Add((time.Hour * 24) * 30).Unix()
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"role":     user.Role,
		"exp":      exp,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err := tokenClaims.SignedString([]byte(s.conf.JWT_SECRET_KEY))

	if err != nil {
		log.Errorf("[Login] err create access token %+v", err)
		return
	}

	resp = &dto.LoginRes{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Username:    user.Username,
		Photo:       user.Photo,
		Phone:       user.Phone,
		AccessToken: AccessToken,
		Role:        user.Role,
		Exp:         exp,
	}
	return
}

func (s *UserService) CreateUser(ctx context.Context, payload dto.PayloadCreateUser, claims *authutil.UserClaims) (resp *model.User, err error) {
	userPayload := model.User{
		Username:   payload.Username,
		Name:       payload.Name,
		Email:      payload.Email,
		NIK:        payload.NIK,
		Prefix:     "+62",
		Photo:      payload.ProfilePath,
		MerchantID: payload.MerchantID,
		Phone:      payload.Email,
		Role:       payload.Role,
	}
	pass, _ := helper.HashPassword(payload.Password)
	userPayload.Password = pass

	resp, err = s.UserRepository.Create(ctx, &userPayload)
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload, claims *authutil.UserClaims) (resp pagination.DefaultPagination, err error) {
	resp, err = s.UserRepository.FindAllUserPaginated(ctx, payload, claims)
	if err != nil {
		log.Errorf("[user.go][GetUserPaginated] err repository at service :%+v", err)
		return
	}

	respToDto := make([]*dto.UserRow, 0)
	list, ok := resp.Items.([]*model.User)
	if ok {
		for _, v := range list {
			temp := &dto.UserRow{
				ID:          v.ID,
				Phone:       v.Phone,
				Name:        v.Name,
				Username:    v.Username,
				NIK:         v.NIK,
				ProfilePath: v.Photo,
				Email:       v.Email,
				MerchantID:  v.MerchantID,
				Role:        v.Role,
			}

			if temp.ProfilePath != nil {
				*temp.ProfilePath = s.conf.AWS_S3_URL + "/" + *temp.ProfilePath
			}
			respToDto = append(respToDto, temp)
		}
	}
	resp.Items = respToDto
	return
}

func (s *UserService) GetById(ctx context.Context, id int) (resp *dto.UserRowDetail, err error) {
	row, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	resp = &dto.UserRowDetail{
		ID:          row.ID,
		Phone:       row.Phone,
		Name:        row.Name,
		Username:    row.Username,
		NIK:         row.NIK,
		ProfilePath: row.Photo,
		Email:       row.Email,
		MerchantID:  row.MerchantID,
		Role:        row.Role,
	}

	if resp.ProfilePath != nil {
		*resp.ProfilePath = s.conf.AWS_S3_URL + "/" + *resp.ProfilePath
	}
	return
}

func (s *UserService) UpdateUser(ctx context.Context, payload dto.PayloadUpdateUser, id int, claims *authutil.UserClaims) (resp *model.User, err error) {
	user, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}
	userPayload := model.User{
		Username:   payload.Username,
		Name:       payload.Name,
		Email:      payload.Email,
		NIK:        payload.NIK,
		Prefix:     "+62",
		Photo:      payload.ProfilePath,
		MerchantID: payload.MerchantID,
		Phone:      payload.Email,
		Role:       payload.Role,
	}

	if payload.Password != nil {
		newPassword, _ := helper.HashPassword(strings.Trim(*payload.Password, " "))
		userPayload.Password = newPassword
	}

	resp, err = s.UserRepository.Update(ctx, &userPayload, int(user.ID))
	if err != nil {
		log.Errorf("[user.go][UpdateUser] err create user :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {
	resp, err = s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) DeleteUserById(ctx context.Context, id int, claims *authutil.UserClaims) (resp *int64, err error) {
	resp, err = s.UserRepository.DeleteUserById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) UpdateUserProfile(ctx context.Context, id int, password string) (resp *int64, err error) {
	newPassword, _ := helper.HashPassword(strings.Trim(password, " "))
	resp, err = s.UserRepository.UpdatePasswordByUserId(ctx, id, &newPassword)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error) {
	resp, err = s.UserRepository.GetUserByIdToken(ctx, userId)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	return
}

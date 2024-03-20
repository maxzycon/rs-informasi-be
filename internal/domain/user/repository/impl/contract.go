package impl

import (
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"gorm.io/gorm"
)

type NewUserRepository struct {
	Conf *config.Config
	Db   *gorm.DB
}
type UserRepository struct {
	conf *config.Config
	db   *gorm.DB
}

func New(params *NewUserRepository) *UserRepository {
	return &UserRepository{
		conf: params.Conf,
		db:   params.Db,
	}
}

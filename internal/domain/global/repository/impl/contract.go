package impl

import (
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NewGlobalRepository struct {
	Conf *config.Config
	Db   *gorm.DB
	Log  *logrus.Logger
}
type GlobalRepository struct {
	conf *config.Config
	db   *gorm.DB
	log  *logrus.Logger
}

func New(params *NewGlobalRepository) *GlobalRepository {
	return &GlobalRepository{
		conf: params.Conf,
		db:   params.Db,
		log:  params.Log,
	}
}

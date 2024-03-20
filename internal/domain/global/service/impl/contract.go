package impl

import (
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/repository"
	S3Service "github.com/maxzycon/rs-informasi-be/internal/domain/s3/service/impl"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NewGlobalServiceParams struct {
	Conf             *config.Config
	GlobalRepository repository.GlobalRepository
	S3Service        *S3Service.S3Service
	Db               *gorm.DB
	Log              *logrus.Logger
}
type GlobalService struct {
	conf             *config.Config
	globalRepository repository.GlobalRepository
	s3Service        *S3Service.S3Service
	db               *gorm.DB
	log              *logrus.Logger
}

func New(params *NewGlobalServiceParams) *GlobalService {
	return &GlobalService{
		conf:             params.Conf,
		globalRepository: params.GlobalRepository,
		s3Service:        params.S3Service,
		db:               params.Db,
		log:              params.Log,
	}
}

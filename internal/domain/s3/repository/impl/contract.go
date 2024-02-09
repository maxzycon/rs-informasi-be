package impl

import (
	"github.com/maxzycon/rs-farmasi-be/internal/config"
	"gorm.io/gorm"
)

type NewS3RepositoryParams struct {
	Conf *config.Config
	Db   *gorm.DB
}

type S3Repository struct {
	conf *config.Config
	db   *gorm.DB
}

func New(params *NewS3RepositoryParams) *S3Repository {
	return &S3Repository{
		conf: params.Conf,
		db:   params.Db,
	}
}

package impl

import (
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/maxzycon/rs-informasi-be/internal/domain/s3/repository"
)

type NewS3ServiceParams struct {
	Conf         *config.Config
	S3Repository repository.S3Repository
}

type S3Service struct {
	conf         *config.Config
	S3Repository repository.S3Repository
}

func New(params *NewS3ServiceParams) *S3Service {
	return &S3Service{
		conf:         params.Conf,
		S3Repository: params.S3Repository,
	}
}

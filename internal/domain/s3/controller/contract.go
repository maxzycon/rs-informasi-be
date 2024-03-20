package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/rs-informasi-be/internal/config"
	"github.com/maxzycon/rs-informasi-be/internal/domain/s3/service"
)

const (
	Create         = "/s3/upload"
	Update         = "/s3/:id"
	Delete         = "/s3/upload"
	GetS3Paginated = "/s3"
)

type S3ControllerParams struct {
	V1        fiber.Router
	Conf      *config.Config
	S3Service service.S3Service
}
type S3Controller struct {
	v1        fiber.Router
	conf      *config.Config
	s3Service service.S3Service
}

func New(params *S3ControllerParams) *S3Controller {
	return &S3Controller{
		v1:        params.V1,
		conf:      params.Conf,
		s3Service: params.S3Service,
	}
}
func (pc *S3Controller) Init() {
	pc.v1.Post(Create, pc.handlerUploadToS3)
	pc.v1.Delete(Delete, pc.handlerDeleteFileS3)
}

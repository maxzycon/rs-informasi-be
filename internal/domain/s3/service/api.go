package service

import (
	"context"
	"mime/multipart"

	"github.com/maxzycon/rs-farmasi-be/internal/domain/s3/dto"
)

type S3Service interface {
	UploadFileToS3(ctx context.Context, file *multipart.FileHeader, bucket string, folder *string) (resp string, err error)
	DeleteFileS3(ctx context.Context, payload *dto.PayloadDeleteS3Path, bucket string, folder *string) (resp bool, err error)
}

package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-farmasi-be/internal/domain/s3/dto"
	"github.com/maxzycon/rs-farmasi-be/pkg/constant"
	"github.com/maxzycon/rs-farmasi-be/pkg/httputil"
)

func (c *S3Controller) handlerUploadToS3(f *fiber.Ctx) (err error) {
	file, err := f.FormFile("file")
	if err != nil {
		log.Errorf("[s3.go][] file err :%+v", err)
		return
	}

	fmt.Println()

	folder := f.Query("folder", "")

	path, err := c.s3Service.UploadFileToS3(f.Context(), file, "go-clinic-bucket", &folder)
	if err != nil {
		log.Errorf("[s3.go][] err save to s3 at controller :%+v", err)
		return
	}

	return httputil.WriteSuccessResponse(f, fiber.Map{
		"path": path,
	})
}

func (c *S3Controller) handlerDeleteFileS3(f *fiber.Ctx) (err error) {
	payload := dto.PayloadDeleteS3Path{}
	err = f.BodyParser(&payload)

	if err != nil {
		log.Errorf("[s3.go][handlerDeleteFileS3] file err :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	folder := f.Query("folder", "")

	_, err = c.s3Service.DeleteFileS3(f.Context(), &payload, "go-clinic-bucket", &folder)
	if err != nil {
		log.Errorf("[s3.go][handlerDeleteFileS3] err save to s3 at controller :%+v", err)
		return
	}

	return httputil.WriteSuccessResponse(f, constant.Success)
}

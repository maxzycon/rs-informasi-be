package impl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/rs-informasi-be/internal/domain/s3/dto"
)

func (s *S3Service) UploadFileToS3(ctx context.Context, file *multipart.FileHeader, bucket string, folder *string) (resp string, err error) {
	ff, err := file.Open()
	if err != nil {
		log.Errorf("[s3.go][UploadFileToS3] err open file :%+v", err)
		return
	}

	defer ff.Close()

	buffer, err := io.ReadAll(ff)

	if err != nil {
		log.Errorf("[s3.go][] err read buffer :%+v", err)
		return
	}

	newFileName := strings.Replace(uuid.New().String(), "-", "", -1) + filepath.Ext(file.Filename)
	key := aws.String(newFileName)

	if strings.Contains(http.DetectContentType(buffer), "image") {
		// if name, converted, err := image.ImageProcessing(buffer, 40); err != nil {
		// 	log.Errorf("[s3.go][] err compress image :%+v", err)
		// 	return "", err
		// } else {
		// 	buffer = converted
		// 	key = aws.String(name)
		// }
	}

	if err != nil {
		log.Errorf("[s3.go][] err read buffer :%+v", err)
		return
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(s.conf.AWS_S3_REGION), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s.conf.AWS_S3_ACCESS_KEY_ID, s.conf.AWS_S3_SECRET_KEY, "")))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := s3.NewFromConfig(cfg)

	if folder != nil && *folder != "" {
		key = aws.String(fmt.Sprintf("%s/%s", *folder, *key))
	}

	// Build the request with its input parameters
	conf := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         key,
		Body:        bytes.NewReader(buffer),
		ACL:         "public-read",
		ContentType: aws.String(http.DetectContentType(buffer)),
	}

	// check image or not
	if strings.Contains(http.DetectContentType(buffer), "image") {
		conf.ContentType = aws.String("image/webp")
	} else {
		conf.Body = bytes.NewReader(buffer)
	}

	_, err = svc.PutObject(ctx, conf)

	if err != nil {
		log.Errorf("[s3.go][UploadFileToS3] err upload file s3 :%+v", err)
		return
	}

	resp = *key
	return
}

func (s *S3Service) UploadBufferExcelToFileS3(ctx context.Context, buffer []byte, bucket string, folder string) (resp string, err error) {
	newFileName := strings.Replace(uuid.New().String(), "-", "", -1) + ".xlsx"
	key := aws.String(newFileName)

	if err != nil {
		log.Errorf("[s3.go][] err read buffer :%+v", err)
		return
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(s.conf.AWS_S3_REGION), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s.conf.AWS_S3_ACCESS_KEY_ID, s.conf.AWS_S3_SECRET_KEY, "")))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := s3.NewFromConfig(cfg)

	if folder != "" {
		key = aws.String(fmt.Sprintf("%s/%s", folder, *key))
	}

	// Build the request with its input parameters
	conf := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         key,
		Body:        bytes.NewReader(buffer),
		ACL:         "public-read",
		ContentType: aws.String(http.DetectContentType(buffer)),
	}

	_, err = svc.PutObject(ctx, conf)

	if err != nil {
		log.Errorf("[s3.go][UploadFileToS3] err upload file s3 :%+v", err)
		return
	}

	resp = *key
	return
}

func (s *S3Service) DeleteFileS3(ctx context.Context, payload *dto.PayloadDeleteS3Path, bucket string, folder *string) (resp bool, err error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(s.conf.AWS_S3_REGION), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s.conf.AWS_S3_ACCESS_KEY_ID, s.conf.AWS_S3_SECRET_KEY, "")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)

	_, err = svc.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(payload.Path),
	})

	if err != nil {
		log.Errorf("[s3.go][UploadFileToS3] err upload file s3 :%+v", err)
		return
	}

	resp = true
	return
}

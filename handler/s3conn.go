package handler

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"os"
)

var log = logrus.New()

func Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Debug()
	}

	src, err := file.Open()
	if err != nil {
		log.Debug(err)
	}
	defer src.Close()
	result, err := UploadToS3(file.Filename, src)
	if err != nil {
		log.Debug(err)
	}
	// result 를 기반으로 UploadResult 구조체에
	// metadata 추가
	// 그리고 그 result 를 기반으로 ffprobe 로 메타데이터를 뽑아오고
	// data 에다가 정보 추가

	resultProbe := FFprobe(result)
	return c.JSON(http.StatusOK, echo.Map{
		"path": result,
		// 일단 format 만 가져오게끔 해놓음..
		"result": resultProbe["format"],
	})
}

func UploadToS3(filename string, src multipart.File) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	s3BucketName := os.Getenv("BUCKET_NAME")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Debug(err)
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		log.Debug(err)
	}
	return result.Location, nil
}

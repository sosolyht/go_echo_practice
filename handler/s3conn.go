package handler

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	result, err := UploadToS3(c, file.Filename, src)
	if err != nil {
		return err
	}
	//fmt.Println("result:", result)
	// result 를 기반으로 UploadResult 구조체에
	// metadata 추가
	// 그리고 그 result 를 기반으로 ffprobe 로 메타데이터를 뽑아오고
	// data 에다가 정보 추가

	a := FFprobe(result)
	return c.JSON(http.StatusOK, echo.Map{
		"path":   result,
		"result": a,
	})
}

func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	s3BucketName := os.Getenv("BUCKET_NAME")

	logger := c.Logger()
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	return result.Location, nil
}

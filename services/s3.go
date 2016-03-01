package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"log"
	"mime/multipart"
	"os"
)

var SVC *s3.S3
var BUCKET *string

func init() {
	godotenv.Load()

	BUCKET = aws.String(os.Getenv("BUCKET"))

	SVC = s3.New(
		session.New(),
		aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithS3ForcePathStyle(true),
	)
}

func TestS3() {
	// LIST OBJECTS
	resp, err := SVC.ListObjects(&s3.ListObjectsInput{
		Bucket: BUCKET,
	})

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	log.Println(resp)
}

func UploadS3(file multipart.File, key string) error {
	_, err := SVC.PutObject(&s3.PutObjectInput{
		Bucket: BUCKET,
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
		Body:   file,
	})

	return err
}

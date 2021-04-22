package s3Service

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/avaldigitallabs/guild-golang/upload-to-s3/constants"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func UploadToBucket(file string, svc s3iface.S3API) error {
	log.Printf("Begin file upload: %v", file)

	text, err := readFile(file)
	if err != nil {
		return err
	}

	body := []byte(text)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(constants.S3_BUCKET_NAME),
		Key:    aws.String(constants.S3_KEY_PREFIX + file),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return err
	}

	log.Printf("Completed file upload: %v", file)

	return nil
}

func DownloadFromBucket(file string, svc s3iface.S3API) (io.ReadCloser, error) {
	log.Printf("Begin file download: %v", file)

	objectOutput, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(constants.S3_BUCKET_NAME),
		Key:    aws.String(constants.S3_KEY_PREFIX + file),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Completed file download: %v", file)

	return objectOutput.Body, nil
}

func readFile(fileLocation string) (string, error) {
	content, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	text := string(content)
	return text, nil
}

package adapter

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
)

func UploadToBucket(file string) error {
	log.Println("Begin file upload...!")

	text, err := readFile(file)
	if err != nil {
		return err
	}

	body := []byte(text)

	sess, err := session.NewSessionWithOptions(session.Options{})
	if err != nil {
		return err
	}

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String("guild-test-bkt"),
		Key:    aws.String("test-file.json"),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return err
	}

	log.Println("File upload completed!")

	return nil
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

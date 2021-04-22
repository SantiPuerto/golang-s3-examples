package business

import (
	"log"

	"github.com/avaldigitallabs/guild-golang/upload-to-s3/constants"
	"github.com/avaldigitallabs/guild-golang/upload-to-s3/s3Service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ExecuteS3Operation(operation, fileName string) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	})
	if err != nil {
		log.Println("Error iniciando la sessión")
	}
	s3ServiceConnection := s3.New(sess)

	switch operation {
	case constants.GET_FILE:
		download, err := s3Service.DownloadFromBucket(fileName, s3ServiceConnection)
		handleError(err)
		log.Printf("Download Result=%v", download)
	case constants.PUT_FILE:
		err := s3Service.UploadToBucket(fileName, s3ServiceConnection)
		handleError(err)
	default:
		println("No option selected")
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal("Error during the execution of S3 operation: ", err)
	}
}

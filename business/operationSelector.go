package business

import (
	"github.com/avaldigitallabs/guild-golang/upload-to-s3/constants"
	"github.com/avaldigitallabs/guild-golang/upload-to-s3/s3Service"
	"log"
)

func ExecuteS3Operation(operation, fileName string) {
	switch operation {
	case constants.GET_FILE:
		download, err := s3Service.DownloadFromBucket(fileName)
		handleError(err)
		log.Printf("Download Result=%v", download)
	case constants.PUT_FILE:
		err := s3Service.UploadToBucket(fileName)
		handleError(err)
	default:
		println("No option selected")
	}
}

func handleError(err error){
	if err != nil {
		log.Fatal("Error during the execution of S3 operation: ", err)
	}
}

package main

import (
	"github.com/avaldigitallabs/guild-golang/upload-to-s3/business"
	"github.com/avaldigitallabs/guild-golang/upload-to-s3/constants"
	"log"
)

func main() {
	log.Println("Starting demo...")

	business.ExecuteS3Operation(constants.PUT_FILE, "test-file.json")

	log.Println("Done!")
}

package adapter

import "log"

func Run() {
	log.Println("Starting...")
	err := UploadToBucket("test-file.json")
	if err != nil {
		log.Fatalf("Error while uploading to S3!", err)
	}
	log.Println("Done.")
}

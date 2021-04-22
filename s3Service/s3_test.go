package s3Service

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if input.Bucket == aws.String("bucket_test") && input.Key == aws.String("object_test") {
		return &s3.GetObjectOutput{
			Body: ioutil.NopCloser(strings.NewReader("hello world")),
		}, nil
	} else {
		return nil, awserr.New("Error", "Error", errors.New("Error"))
	}
}

func TestS3Service(t *testing.T) {
	mockSvc := &mockS3Client{}
	_, err := DownloadFromBucket("test-file.json", mockSvc)
	if err != nil {
		assert.Errorf(t, err, "Error")
	}
}

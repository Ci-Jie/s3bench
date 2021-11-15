package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadInput ...
type UploadInput struct {
	Bucket  string
	Key     string
	Content []byte
}

// UploadOutput ...
type UploadOutput struct {
}

// Upload ...
func (h *Handler) Upload(input *UploadInput) (output *UploadOutput, err error) {
	uploader := s3manager.NewUploader(h.sess)
	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
		Body:   bytes.NewReader(input.Content),
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

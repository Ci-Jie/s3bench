package s3

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

// DeleteObjectInput ...
type DeleteObjectInput struct {
	BucketName string
	Key        string
}

// DeleteObjectOutput ...
type DeleteObjectOutput struct {
}

// DeleteObject ...
func (h *Handler) DeleteObject(input *DeleteObjectInput) (output *DeleteObjectOutput, err error) {
	if _, err = awss3.New(h.sess).DeleteObject(&awss3.DeleteObjectInput{
		Bucket: &input.BucketName,
		Key:    &input.Key,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

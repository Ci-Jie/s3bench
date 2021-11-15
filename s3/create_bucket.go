package s3

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

// CreateBucketInput ...
type CreateBucketInput struct {
	BucketName string
}

// CreateBucketOutput ...
type CreateBucketOutput struct {
}

// CreateBucket ...
func (h *Handler) CreateBucket(input *CreateBucketInput) (output *CreateBucketOutput, err error) {
	if _, err = awss3.New(h.sess).CreateBucket(&awss3.CreateBucketInput{
		Bucket: &input.BucketName,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

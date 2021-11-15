package s3

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

// DeleteBucketInput ...
type DeleteBucketInput struct {
	BucketName string
}

// DeleteBucketOutput ...
type DeleteBucketOutput struct {
}

// DeleteBucket ...
func (h *Handler) DeleteBucket(input *DeleteBucketInput) (output *DeleteBucketOutput, err error) {
	if _, err = awss3.New(h.sess).DeleteBucket(&awss3.DeleteBucketInput{
		Bucket: &input.BucketName,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

package s3

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

// GetObjectInput ...
type GetObjectInput struct {
	BucketName string
	Key        string
}

// GetObjectOutput ...
type GetObjectOutput struct {
	Size int
}

// GetObject ...
func (h *Handler) GetObject(input *GetObjectInput) (output *GetObjectOutput, err error) {
	resp, err := awss3.New(h.sess).GetObject(&awss3.GetObjectInput{
		Bucket: &input.BucketName,
		Key:    &input.Key,
	})
	if err != nil {
		return nil, err
	}
	return &GetObjectOutput{
		Size: int(*resp.ContentLength),
	}, nil
}

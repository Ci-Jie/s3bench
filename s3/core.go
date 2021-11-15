package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Handler ...
type Handler struct {
	endpoint string
	access   string
	secret   string
	region   string
	sess     *session.Session
}

// New ...
func New(endpoint, access, secret, region string) (handler *Handler) {
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			access,
			secret,
			"",
		),
		S3ForcePathStyle: aws.Bool(true)},
	)
	return &Handler{
		endpoint: endpoint,
		access:   access,
		secret:   secret,
		region:   region,
		sess:     sess,
	}
}

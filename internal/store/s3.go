package store

import (
	"context"

	c "github.com/anvidev/nit/config"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewAwsS3Bucket(cfg *c.Config) (*s3.Client, error) {
	creds := credentials.NewStaticCredentialsProvider(cfg.AwsKey, cfg.AwsSecret, "")

	s3Cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(cfg.AwsRegion))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(s3Cfg)

	return client, nil
}

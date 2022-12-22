package cloud

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Action struct {
	*s3.Client
}

func (s *S3Action) NewClient() (*S3Action, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("S3 Client initialization failure")
	}
	s3Client := s3.NewFromConfig(cfg)
	return &S3Action{s3Client}, err
}

func (s *S3Action) GetBucketTags(bucketName string) (*s3.GetBucketTaggingOutput, error) {
	tagOutput, _ := s.Client.GetBucketTagging(context.TODO(),
		&s3.GetBucketTaggingInput{Bucket: aws.String(bucketName)})
	return tagOutput, nil
}

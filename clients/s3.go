package clients

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	Svc *s3.Client
}

// NewS3Client creates a new client for the S3 service.
func NewS3Client(region string) (*Client, error) {
	// Initialize client configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(cfg)

	// Create and return a new client
	return &Client{
		Svc: s3Client,
	}, nil
}

// ListBuckets lists all buckets in the account.
func (c *Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return c.Svc.ListBuckets(ctx, params)
}

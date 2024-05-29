package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"netflix.com/chunker/cmd/nfx-file-chunker/config"
)

var client *s3.Client

func ConfigureS3Client(cfg *config.Config) {
	customResolver := newCustomResolver(cfg)

	// Подгружаем конфигрурацию из ~/.storage/*
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client = s3.NewFromConfig(awsCfg)
}

func newCustomResolver(cfg *config.Config) aws.EndpointResolverWithOptionsFunc {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == cfg.S3.Region {
			return aws.Endpoint{
				PartitionID:   cfg.S3.PartitionId,
				URL:           cfg.S3.Url,
				SigningRegion: cfg.S3.Region,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	return customResolver
}

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/u2takey/ffmpeg-go"
	"log"
	"netflix.com/chunker/cmd/nfx-file-chunker/config"
	"os"
)

func main() {
	//err := ffmpeg.Input("./data/invs.mp4").
	//	Output("./data/out%d.mp4",
	//		ffmpeg.KwArgs{"f": "segment", "c": "copy", "segment_time": "10", "reset_timestamps": "1"}).
	//	OverWriteOutput().ErrorToStdOut().Run()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

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

	// Подгружаем конфигрурацию из ~/.aws/*
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(awsCfg)

	// Запрашиваем список бакетов
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range result.Buckets {
		log.Printf("bucket=%s creation time=%s", aws.ToString(bucket.Name), bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}

	file, err := os.Open("data/invs.mp4")
	if err != nil {
		log.Fatal(err)
	}

	output, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      result.Buckets[0].Name,
		Key:         aws.String("film"),
		Body:        file,
		ContentType: aws.String("mp4"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)

}

package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"path"
)

func UploadContentFiles(contentDir string) (string, error) {
	//Создаю бакет
	bucketName := uuid.New().String()
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	//Переношу файлы в бакет
	files, _ := os.ReadDir(contentDir)
	var eg errgroup.Group
	for _, file := range files {
		file := file
		eg.Go(func() error {
			filePath := path.Join(contentDir, file.Name())
			err := UploadFile(bucketName, file.Name(), filePath)
			if err != nil {
				log.Fatal(err)
				return err
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return "", err
	}
	os.RemoveAll(contentDir)
	return bucketName, nil
}

func UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}
	}
	return err
}

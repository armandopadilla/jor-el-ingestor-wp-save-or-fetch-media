package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

/**
 * Fetch the posts and save to SQS
 **/
func saveToS3(key string, payload string) {

	// Set up the session and client for use.
	s3Session, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	s3Client := s3.New(s3Session)

	buffer := []byte(payload)

	fmt.Println(key)
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("jor-el-data-lake/data-raw-blogs/"),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	return
}

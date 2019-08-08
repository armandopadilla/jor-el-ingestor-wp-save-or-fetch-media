package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Set up the session and client for use.
var sqsSession, _ = session.NewSession(&aws.Config{
	Region:      aws.String(region),
	Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
})

var sqsClient = sqs.New(sqsSession)

/**
 * Add entry into the media Q
 **
func saveToMediaSQS(payload string) {
	if payload == "" {
		fmt.Println("Error", "Payload can not be empty")
		return
	}

	qURL := sqsMediaQURL
	results, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(payload),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}
}*/

/**
* Save the structured data into SQS
***/
func fetchDataFromWPPostQueue() []*sqs.Message {
	fmt.Println("Fetch posts....")

	qURL := sqsWPPostQURL
	result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(10),
	})

	if err != nil { // resp is now filled
		fmt.Println("Error", err)
	}

	if len(result.Messages) == 0 {
		fmt.Println("No posts found")
	}

	return result.Messages
}

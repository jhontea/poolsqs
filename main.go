package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("SQS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("SQS_KEY"), os.Getenv("SQS_SECRET"), ""),
	})
	if err != nil {
		log.Fatalf("New session failed: %s", err)
		return
	}

	svc = sqs.New(sess)
}

var svc *sqs.SQS

func main() {
	chnMessages := make(chan *sqs.Message)

	go pollSqs(chnMessages)

	log.Printf("Listening on stack queue: %s", os.Getenv("SQS_QUEUE_URL"))

	for message := range chnMessages {
		fmt.Printf("%+v\n", message)
		fmt.Println()
	}
}

func pollSqs(chn chan<- *sqs.Message) {

	for {
		queueUrl := os.Getenv("SQS_QUEUE_URL")
		output, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            &queueUrl,
			MaxNumberOfMessages: aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(1),
		})

		if err != nil {
			log.Printf("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {
			chn <- message
		}
	}
}

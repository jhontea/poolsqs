package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Connection struct {
	Key    string
	Secret string
	Region string
}

func NewConnection(region string, key string, secret string) *Connection {
	return &Connection{
		Key:    key,
		Secret: secret,
		Region: region,
	}
}

func (conn *Connection) Connect() (*sqs.SQS, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conn.Region),
		Credentials: credentials.NewStaticCredentials(conn.Key, conn.Secret, ""),
	})
	if err != nil {
		log.Fatalf("New session failed: %s", err)
		return &sqs.SQS{}, err
	}

	svc = sqs.New(sess)

	return svc, nil
}

package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type consumer struct {
	workerPool  int
	workerCount int
	maxMessages int
}

func (c *consumer) Consume() {
	c.workerPool = 50
	c.workerCount = 0
	c.maxMessages = 10

	for w := 1; w <= c.workerPool; w++ {
		go c.worker(w)
	}
}

func (c *consumer) worker(id int) {
	url := "https://sqs.ap-southeast-1.amazonaws.com/656668151939/referral"
	var output *sqs.ReceiveMessageOutput
	var err error

	for {
		output, err = ReceiveMessage(&url)
		if err != nil {
			// log error
			fmt.Println(err.Error())
		}
	}

	var wg sync.WaitGroup

	for _, m := range output.Messages {
		wg.Add(1)

		go func(m *sqs.Message) {
			defer wg.Done()

			fmt.Printf("%v", m)
			fmt.Println()
		}(m)
	}

	wg.Wait()
}

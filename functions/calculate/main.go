package main

import (
	"context"
	"encoding/json"
	"hello-sqs/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var log = logger.New("main")

// Body of the SQS Event message
type Body struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type Resp struct {
	Total int `json:"total"`
}

// possible handler functions: https://pkg.go.dev/github.com/aws/aws-lambda-go@v1.41.0/lambda#Start
func handler(ctx context.Context, snsEvent events.SNSEvent) (resp Resp, err error) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		log.Infof("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

		body := Body{}
		if err = json.Unmarshal([]byte(snsRecord.Message), &body); err != nil {
			return
		}
		sum := body.FirstNumber + body.SecondNumber
		log.Infof("Calculation finished. %d + %d = %d", body.FirstNumber, body.SecondNumber, sum)
		resp.Total = sum
	}

	return resp, nil
}

func main() {
	lambda.Start(handler)
}

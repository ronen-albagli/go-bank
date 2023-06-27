package consumers

import (
	"bank/app/adapter"
	"bank/config"
	"bank/usecase"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type QuotaType struct {
	Credits int32
	Seats   int32
}

type TransactionInput struct {
	AccountId int64
	Quota     QuotaType
	Reason    string
}

type NewAssetResponse struct {
	TransactionId string `json:"transactionId"`
}

func InvoiceConsumer() {
	var parsedInput TransactionInput
	conf := new(config.Config)
	conf.SetConf()

	sqsClient, err := conf.GetAWSSqsClient()

	queueUrl := "https://sqs.us-east-1.amazonaws.com/020761065253/invoice_completed.fifo"

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(1), // Maximum number of messages to retrieve
		WaitTimeSeconds:     aws.Int64(2), // Long polling wait time (in seconds)
	}

	if err != nil {
		panic(err)
	}

	for {
		// Receive messages from the queue
		resp, err := sqsClient.ReceiveMessage(params)
		if err != nil {
			fmt.Println("Failed to receive messages:", err)
			continue
		}

		// Process the received messages
		for _, msg := range resp.Messages {
			// Extract the message body
			body := aws.StringValue(msg.Body)

			// Perform your desired operations with the message body
			fmt.Println("Received message:", body)

			json.Unmarshal([]byte(body), &parsedInput)

			mongoColl := conf.GetLedgerGateway()

			configuration := usecase.Config{
				LedgerCollection: &adapter.LedgerMongoGateway{
					Collection: *mongoColl,
				},
			}

			var input = &usecase.AssetInputTest{}

			input.EventName = "GRANT"
			input.Amount = int64(parsedInput.Quota.Credits)
			input.AssetType = "USD"
			input.Reason = "Payment_created"
			input.AccountId = parsedInput.AccountId

			usecase := &usecase.AddCreditUseCaseStruct{
				Config: configuration,
			}

			transactionId, error1 := usecase.Do(*input)

			if error1 != nil {
				fmt.Println(error1)
				return
			}

			var response = &NewAssetResponse{
				TransactionId: transactionId,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(jsonResponse)

			// Delete the message from the queue
			_, deleteREr := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueUrl),
				ReceiptHandle: msg.ReceiptHandle,
			})

			if deleteREr != nil {
				fmt.Println("Failed to delete message:", err)
			}
		}
	}
}

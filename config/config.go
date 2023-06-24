package config

import (
	"context"
	"fmt"
	"os"
	"time"

	adapter "bank/app/adapter"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo interface{}

type SetConf interface{}
type GetMongoClient interface{}

// type GetLedgerGateway interface{}

type Config struct {
	mongo *mongo.Client
	// envList *Conf
}

type Conf struct {
	MongoUri  string `yaml:"mongoUri"`
	AwsAccess string `yaml:"awsAccessKey"`
	AwsSecret string `yaml:"awsSecretKey"`
}

type GetLedgerGateway struct {
	LedgerCollecton *mongo.Collection
}

// type getAWSSqsClient struct {
// 	*sqs.ReceiveMessageOutput
// }

func (c *Config) GetMongoInMemory() *adapter.LedgerMongoInMemoryGateway {
	// client, err := c.GetMongoClient()

	mongoMemory := &adapter.LedgerMongoInMemoryGateway{}
	// if err != nil {
	// 	return nil
	// }

	// ledgerGateway := client.Database("Mongo").Collection("ledger")

	// if err != nil {
	// 	return nil
	// }

	return mongoMemory
}

func (c *Config) GetLedgerGateway() *mongo.Collection {
	client, err := c.GetMongoClient()

	if err != nil {
		return nil
	}

	ledgerGateway := client.Database("Mongo").Collection("ledger")

	if err != nil {
		return nil
	}

	return ledgerGateway
}

func (c *Config) GetAWSSqsClient() (*sqs.SQS, error) {
	awsAccess := os.Getenv("AWS_ACCESS_KEY")
	awsSecret := os.Getenv("AWS_SECRET_KEY")

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsAccess, awsSecret, ""),
		Region:      aws.String("us-east-1"), // Replace with your desired region
	})

	if err != nil {
		fmt.Println("Failed to create session:", err)
		return nil, err
	}

	// Create an SQS service client
	svc := sqs.New(sess)

	// Specify the URL of the queue to consume messages from
	return svc, nil
}

// func (c *Config) GetInvoiceQueueGateway() (*sqs.ReceiveMessageOutput, error) {
// 	client, err := c.getAWSSqsClient()

// 	if err != nil {
// 		panic(err)
// 	}

// 	params := &sqs.ReceiveMessageInput{
// 		QueueUrl:            aws.String("https://sqs.us-east-1.amazonaws.com/020761065253/invoice_completed.fifo"),
// 		MaxNumberOfMessages: aws.Int64(1), // Maximum number of messages to retrieve
// 		WaitTimeSeconds:     aws.Int64(2), // Long polling wait time (in seconds)
// 	}

// 	return client,

// params := client.R{
// 	QueueUrl:            aws.String(queueURL),
// 	MaxNumberOfMessages: aws.Int64(10), // Maximum number of messages to retrieve
// 	WaitTimeSeconds:     aws.Int64(20), // Long polling wait time (in seconds)
// }
// }

func (c *Config) GetMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(900)*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.getMongoURI()))

	c.mongo = client
	if err != nil {
		println("NO CLIENT")
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Mongo is Connected")
	return c.mongo, nil
}

func (c *Config) getMongoURI() string {
	return os.Getenv("MONGO_CON_STR")
}

func (c *Config) SetConf() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// c.envList = cfg
}

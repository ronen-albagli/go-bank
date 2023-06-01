package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

type Mongo interface{}

type SetConf interface{}
type GetMongoClient interface{}

// type GetLedgerGateway interface{}

type Config struct {
	mongo   *mongo.Client
	envList *Conf
}

type Conf struct {
	MongoUri string `yaml:"mongoUri"`
}

type GetLedgerGateway struct {
	LedgerCollecton *mongo.Collection
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

func (c *Config) GetMongoClient() (*mongo.Client, error) {
	println(c.getMongoURI())
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
	return c.envList.MongoUri
}

func (c *Config) SetConf() {
	if c.envList == nil {
		file, err := os.Open("config.yaml")
		if err != nil {

			panic(err)
		}
		defer file.Close()
		cfg := &Conf{}
		yd := yaml.NewDecoder(file)
		err = yd.Decode(cfg)

		if err != nil {
			panic(err)
		}

		c.envList = cfg
	}
}

package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	types "bank/app/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LedgerEvent struct {
	amount    int64
	assetType string
	eventName string
	reason    string
}

type LedgerMongoGateway struct {
	mongo.Collection
}

func (m *LedgerMongoGateway) Store(input types.NewAssetInput) error {
	ctx := context.Background()

	fmt.Println("Saved in DB", input)

	data, err := m.InsertOne(
		ctx,
		bson.M{
			"amount":    input.Amount,
			"eventType": "GRANT",
			"event":     input.EventName,
			"timestamp": time.Now(),
			"assetType": input.AssetType,
			"reason":    input.Reason,
		})

	if err != nil {
		return errors.New("failed to insert to ledger")
	}

	fmt.Println(data)
	return nil
}

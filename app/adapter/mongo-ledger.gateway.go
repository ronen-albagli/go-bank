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

// type LedgerEvent struct {
// 	AccountId     int64
// 	Amount        int64
// 	AssetType     string
// 	EventName     string
// 	Reason        string
// 	TransactionId string
// 	Version       int64
// }

type LedgerMongoGateway struct {
	mongo.Collection
}

func (m *LedgerMongoGateway) Store(input types.NewAssetInput) error {
	ctx := context.Background()

	fmt.Println("Saved in DB", input)

	data, err := m.InsertOne(
		ctx,
		bson.M{
			"accountId":     input.AccountId,
			"amount":        input.Amount,
			"eventType":     "GRANT",
			"event":         input.EventName,
			"timestamp":     time.Now(),
			"assetType":     input.AssetType,
			"reason":        input.Reason,
			"transactionId": input.TransactionId,
			"version":       input.Version,
		})

	if err != nil {
		return errors.New("failed to insert to ledger")
	}

	fmt.Println(data)
	return nil
}

func (m *LedgerMongoGateway) InitLedger(accountId int64) types.LedgerEvent {
	var cur *mongo.Cursor
	var err error
	ctx := context.Background()
	var doc types.LedgerEvent
	var ledgerEvents []types.LedgerEvent

	filter := bson.D{{Key: "accountId", Value: accountId}}

	if cur, err = m.Find(ctx, filter); err != nil {
		panic(err)
	}

	cur.Decode(&doc)

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		cur.Decode(&doc)
		ledgerEvents = append(ledgerEvents, doc)
	}

	return doc

}

package adapter

import (
	"context"
	"errors"
	"fmt"

	entities "bank/app/entities"
	types "bank/app/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LedgerMongoGateway struct {
	mongo.Collection
}

func (m *LedgerMongoGateway) Store(LedgerEvents []types.LedgerEvent) error {
	ctx := context.Background()

	for _, LedgerEvent := range LedgerEvents {
		data, err := m.InsertOne(
			ctx,
			bson.M{
				"accountId":     LedgerEvent.AccountId,
				"amount":        LedgerEvent.Amount,
				"eventName":     LedgerEvent.EventName,
				"timestamp":     LedgerEvent.Timestamp,
				"assetType":     LedgerEvent.AssetType,
				"reason":        LedgerEvent.Reason,
				"transactionId": LedgerEvent.TransactionId,
				"version":       LedgerEvent.Version,
			})

		if err != nil {
			return errors.New("failed to insert to ledger")
		}

		fmt.Println(data)
	}

	return nil
}

func (m *LedgerMongoGateway) InitLedger(accountId int64) *entities.LedgerEntity {
	var cur *mongo.Cursor
	var err error
	ctx := context.Background()
	var ledgerEvents []types.LedgerEvent

	filter := bson.D{{Key: "accountId", Value: accountId}}

	if cur, err = m.Find(ctx, filter); err != nil {
		panic(err)
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &ledgerEvents); err != nil {
		panic(err)
	}

	ledger := &entities.LedgerEntity{
		AccountId: accountId,
		Version:   0,
	}

	ledger.ApplyEvents(ledgerEvents)

	return ledger

}

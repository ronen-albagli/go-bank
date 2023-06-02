package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type LedgerGrantInput struct {
}

type INewAssetInput interface {
	NewAssetInput
}

type NewAssetInput struct {
	AccountId     int64
	EventName     string
	Amount        int64
	AssetType     string
	Reason        string
	TransactionId string
	Version       int64
	Timestamp     int64
}

type LedgerEvent struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	AccountId     int64              `bson:"accountId,omitempty"`
	Amount        int64              `bson:"amount,omitempty"`
	AssetType     string             `bson:"assetType,omitempty"`
	EventName     string             `bson:"eventName,omitempty"`
	Reason        string             `bson:"reason,omitempty"`
	TransactionId string             `bson:"transactionId,omitempty"`
	Version       int64              `bson:"version,omitempty"`
	Timestamp     int64              `bson:"timestamp,omitempty"`
}

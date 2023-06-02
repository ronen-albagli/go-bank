package types

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
}

type LedgerEvent struct {
	AccountId     int64
	Amount        int64
	AssetType     string
	EventName     string
	Reason        string
	TransactionId string
	Version       int64
}

package types

type LedgerGrantInput struct {
}

type INewAssetInput interface {
	NewAssetInput
}

type NewAssetInput struct {
	EventName string
	Amount    int64
	AssetType string
	Reason    string
}

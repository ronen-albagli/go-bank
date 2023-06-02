package ports

import (
	types "bank/app/types"
)

type AssetInput struct {
}

type LedgerGateway interface {
	Store(asset types.NewAssetInput) error
	InitLedger(accountId int64) types.LedgerEvent
}

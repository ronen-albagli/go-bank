package ports

import (
	types "bank/app/types"
)

type AssetInput struct {
}

type LedgerGateway interface {
	Store(asset types.NewAssetInput) error
}

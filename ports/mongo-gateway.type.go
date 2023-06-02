package ports

import (
	entities "bank/app/entities"
	types "bank/app/types"
)

type AssetInput struct {
}

type LedgerGateway interface {
	Store(asset []types.LedgerEvent) error
	InitLedger(accountId int64) *entities.LedgerEntity
}

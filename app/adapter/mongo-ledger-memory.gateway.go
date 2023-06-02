package adapter

import (
	entities "bank/app/entities"
	types "bank/app/types"
	"bank/gateway"
	"fmt"
)

type LedgerMongoInMemoryGateway struct {
	Collection interface{}
}

func (m *LedgerMongoInMemoryGateway) Store(asset []types.LedgerEvent) error {
	fmt.Println("Saved in memory")

	gateway.Save()

	return nil
}

func (m *LedgerMongoInMemoryGateway) InitLedger(accountId int64) *entities.LedgerEntity {
	var ledgerEvents []types.LedgerEvent

	ledger := &entities.LedgerEntity{
		AccountId: accountId,
		Version:   0,
	}

	ledger.ApplyEvents(ledgerEvents)

	return ledger
}

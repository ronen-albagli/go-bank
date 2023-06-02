package adapter

import (
	types "bank/app/types"
	"bank/gateway"
	"bank/ports"
	"fmt"
)

type LedgerMongoInMemoryGateway struct {
	ports.LedgerGateway
}

func (m *LedgerMongoInMemoryGateway) Store(asset []types.LedgerEvent) error {
	fmt.Println("Saved in memory")

	gateway.Save()

	return nil
}

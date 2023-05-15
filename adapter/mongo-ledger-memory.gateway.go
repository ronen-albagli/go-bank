package adapter

import (
	"bank/gateway"
	"fmt"
)

type LedgerGateway interface {
	Store(asset Input)
}

type LedgerMongoInMemoryGateway struct{}

func (m *LedgerMongoInMemoryGateway) Store(asset Input) {
	fmt.Println("Saved in memory")

	gateway.Save()
}

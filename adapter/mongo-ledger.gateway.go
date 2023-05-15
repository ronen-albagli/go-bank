package adapter

import (
	"bank/gateway"
	"fmt"
)

type Input struct {
	eventName string
	amount    int64
	assetType string
	reason    string
}

// type LedgerGateway interface {
// 	Store(asset Input)
// }

type LedgerMongoGateway struct{}

func (m *LedgerMongoGateway) Store(asset Input) {
	fmt.Println("Saved in DB")
	// m.save(asset)

	gateway.Save()
}

// func GetGateway() *LedgerMemoryGateway {
// 	var gateway *LedgerMemoryGateway = new(LedgerMemoryGateway)

// 	return gateway
// }

// type ILedgerMemoryGateway interface {
// 	save(a Asset)
// }

// type LedgerMemoryGateway struct{}

// func (m *LedgerMemoryGateway) save(asset Asset) {
// 	fmt.Println("Saved in memory")

// 	m.save(asset)
// }

// func (mo interface{ LedgerGateway }) save(m LedgerMemoryGateway, asset Asset) {
// 	fmt.Println("Saved in memory")
// 	mo.save(asset)
// }

package usecase

import (
	adapter "bank/adapter"
	"fmt"
)

type AddCreditInput struct {
	amount    int64
	assetType string
	reason    string
}

type Input1 struct {
	eventName string
	amount    int64
	assetType string
	reason    string
}

type IAddCreditUseCase interface {
	do(asset Input1)
}

// func (ia AddCreditUseCase) AddCreditUseCaseStruct() {
// }

type IConfig interface {
}

type Config struct {
	mongo adapter.LedgerGateway
}

type AddCreditUseCaseStruct struct {
	c IAddCreditUseCase
}

func (c *Config) do(asset Input1) {
	a := adapter.Input{}
	c.mongo.Store(a)
	fmt.Println(asset)
}

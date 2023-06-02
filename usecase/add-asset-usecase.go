package usecase

import (
	"bank/ports"
	"fmt"
)

type AssetInputTest struct {
	EventName string
	Amount    int64
	AssetType string
	Reason    string
	AccountId int64
}

type IAddCreditUseCase interface {
	Do(asset AssetInputTest) (string, error)
}

type IConfig interface {
	Config
}

type Config struct {
	LedgerCollection ports.LedgerGateway
}

type AddCreditUseCaseStruct struct {
	IAddCreditUseCase
}

func (c *Config) Do(asset AssetInputTest) (string, error) {
	ledger := c.LedgerCollection.InitLedger(asset.AccountId)

	transactionId, _ := ledger.AddQuota(asset.AssetType, asset.Amount, asset.Reason)

	err := c.LedgerCollection.Store(ledger.GetEvents())

	if err != nil {
		return "", err
	}

	fmt.Println(asset)

	return transactionId, nil
}

package usecase

import (
	types "bank/app/types"
	adapter "bank/ports"
	"fmt"

	"github.com/google/uuid"
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
	LedgerCollection adapter.LedgerGateway
}

type AddCreditUseCaseStruct struct {
	IAddCreditUseCase
}

func (c *Config) Do(asset AssetInputTest) (string, error) {

	ledgerEvents := c.LedgerCollection.InitLedger(1)

	fmt.Println(ledgerEvents)

	transactionId := uuid.New().String()

	a := types.NewAssetInput{}
	a.AccountId = asset.AccountId
	a.Amount = asset.Amount
	a.AssetType = asset.AssetType
	a.Reason = asset.Reason
	a.EventName = "GRANT"
	a.TransactionId = transactionId

	err := c.LedgerCollection.Store(a)

	if err != nil {
		return "", err
	}

	fmt.Println(asset)

	return transactionId, nil
}

package usecase

import (
	"fmt"
)

type AssetInputTest struct {
	EventName string
	Amount    int64
	AssetType string
	Reason    string
	AccountId int64
}

func (addUseCase AddCreditUseCaseStruct) Do(asset AssetInputTest) (string, error) {
	c := addUseCase.Config

	ledger := c.LedgerCollection.InitLedger(asset.AccountId)

	transactionId, _ := ledger.AddQuota(asset.AssetType, asset.Amount, asset.Reason)

	err := c.LedgerCollection.Store(ledger.GetEvents())

	if err != nil {
		return "", err
	}

	fmt.Println(asset)

	return transactionId, nil
}

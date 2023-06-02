package usecase

import (
	"fmt"
)

type IRemoveCreditUseCase interface {
	Do(asset AssetInputTest) (string, error)
}

type RemoveCreditUseCaseStruct struct {
	IAddCreditUseCase
	Config
}

func (removeUseCase RemoveCreditUseCaseStruct) Do(asset AssetInputTest) (string, error) {
	c := removeUseCase.Config
	ledger := c.LedgerCollection.InitLedger(asset.AccountId)

	transactionId, _ := ledger.ReduceQuota(asset.AssetType, asset.Amount, asset.Reason)

	err := c.LedgerCollection.Store(ledger.GetEvents())

	if err != nil {
		return "", err
	}

	fmt.Println(asset)

	return transactionId, nil
}

package usecase

import (
	"fmt"
)

type IUseCreditUseCase interface {
	Do(asset AssetInputTest) (string, error)
}

type UseCreditUseCaseStruct struct {
	IAddCreditUseCase
	Config
}

func (removeUseCase UseCreditUseCaseStruct) Do(asset AssetInputTest) (string, error) {
	c := removeUseCase.Config
	ledger := c.LedgerCollection.InitLedger(asset.AccountId)

	transactionId, useError := ledger.UseQuota(asset.AssetType, asset.Amount, asset.Reason)

	if useError != nil {
		return "", useError
	}

	err := c.LedgerCollection.Store(ledger.GetEvents())

	if err != nil {
		return "", err
	}

	fmt.Println(asset)

	return transactionId, nil
}

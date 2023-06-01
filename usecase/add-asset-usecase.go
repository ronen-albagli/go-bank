package usecase

import (
	types "bank/app/types"
	adapter "bank/ports"
	"fmt"
)

type AssetInputTest struct {
	EventName string
	Amount    int64
	AssetType string
	Reason    string
}

type IAddCreditUseCase interface {
	Do(asset AssetInputTest) error
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

func (c *Config) Do(asset AssetInputTest) error {
	a := types.NewAssetInput{}
	a.Amount = asset.Amount
	a.AssetType = asset.AssetType
	a.Reason = asset.Reason
	a.EventName = "GRANT"

	c.LedgerCollection.Store(a)

	fmt.Println(asset)

	return nil
}

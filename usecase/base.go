package usecase

import (
	"bank/ports"
)

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
	Config
}

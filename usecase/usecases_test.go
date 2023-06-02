package usecase

import (
	adapter "bank/app/adapter"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddLedgerEvent struct {
	ledgerCollection *mongo.Collection
}

func TestAddAssets(t *testing.T) {
	var input = &AssetInputTest{}

	configuration := &Config{
		LedgerCollection: &adapter.LedgerMongoInMemoryGateway{},
	}

	input.EventName = "zazi"
	input.Amount = 10
	input.AssetType = "Shekel"
	input.Reason = "Salary"

	usecase := &AddCreditUseCaseStruct{
		configuration,
	}

	transactionId, err := usecase.Do(*input)

	if err != nil {
		t.Error("Use case failed, Error: ", err)
	}

	if transactionId != "" {
		t.Log("Test pass")
	}

}

func TestHelloEmpty(t *testing.T) {

}

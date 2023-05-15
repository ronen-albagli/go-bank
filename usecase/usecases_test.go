package usecase

import (
	"bank/adapter"
	"testing"
)

func TestAddAssets(t *testing.T) {

	var input = &Input1{}

	config := &Config{
		mongo: &adapter.LedgerMongoGateway{},
	}

	input.amount = 10
	input.assetType = "Shekel"
	input.reason = "Salary"
	input.eventName = "zazi"

	usecase := &AddCreditUseCaseStruct{
		config,
	}

	usecase.c.do(*input)
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	// msg, err := Hello("")
	// if msg != "" || err == nil {
	// 	t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	// }
}

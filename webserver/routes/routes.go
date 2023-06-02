package router

import (
	"encoding/json"
	"net/http"

	adapter "bank/app/adapter"
	"bank/config"
	usecase "bank/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	AccountId int64
	AssetType string
	Amount    int64
	Reason    string
}

type NewAssetResponse struct {
	TransactionId string `json:"transactionId"`
}

func CreateServerRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, "{}")
	})

	app.POST("/transaction/use", func(c *gin.Context) {
		var parsedInput TransactionInput

		if err := c.BindJSON(&parsedInput); err != nil {
			panic("Failed to parse route input")
		}

		conf := new(config.Config)
		conf.SetConf()

		mongoColl := conf.GetLedgerGateway()

		configuration := usecase.Config{
			LedgerCollection: &adapter.LedgerMongoGateway{
				Collection: *mongoColl,
			},
		}

		var input = &usecase.AssetInputTest{}

		input.EventName = "USE"
		input.Amount = parsedInput.Amount
		input.AssetType = parsedInput.AssetType
		input.Reason = parsedInput.Reason
		input.AccountId = parsedInput.AccountId

		usecase := &usecase.UseCreditUseCaseStruct{
			Config: configuration,
		}

		transactionId, error1 := usecase.Do(*input)

		if error1 != nil {
			c.JSON(http.StatusBadRequest, error1.Error())
			return
		}

		var response = &NewAssetResponse{
			TransactionId: transactionId,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusCreated, string(jsonResponse))
	})

	app.POST("/transaction/deduct", func(c *gin.Context) {
		var parsedInput TransactionInput

		if err := c.BindJSON(&parsedInput); err != nil {
			panic("Failed to parse route input")
		}

		conf := new(config.Config)
		conf.SetConf()

		mongoColl := conf.GetLedgerGateway()

		configuration := usecase.Config{
			LedgerCollection: &adapter.LedgerMongoGateway{
				Collection: *mongoColl,
			},
		}

		var input = &usecase.AssetInputTest{}

		input.EventName = "DEDUCT"
		input.Amount = parsedInput.Amount
		input.AssetType = parsedInput.AssetType
		input.Reason = parsedInput.Reason
		input.AccountId = parsedInput.AccountId

		usecase := &usecase.RemoveCreditUseCaseStruct{
			Config: configuration,
		}

		transactionId, err := usecase.Do(*input)

		if err != nil {
			panic("Failed to execute use case")
		}

		var response = &NewAssetResponse{
			TransactionId: transactionId,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusCreated, string(jsonResponse))
	})

	app.POST("/transaction/add", func(c *gin.Context) {
		var parsedInput TransactionInput

		if err := c.BindJSON(&parsedInput); err != nil {
			panic("Failed to parse route input")
		}

		conf := new(config.Config)
		conf.SetConf()

		mongoColl := conf.GetLedgerGateway()

		configuration := usecase.Config{
			LedgerCollection: &adapter.LedgerMongoGateway{
				Collection: *mongoColl,
			},
		}

		var input = &usecase.AssetInputTest{}

		input.EventName = "GRANT"
		input.Amount = parsedInput.Amount
		input.AssetType = parsedInput.AssetType
		input.Reason = "Salary"
		input.AccountId = parsedInput.AccountId

		usecase := &usecase.AddCreditUseCaseStruct{
			Config: configuration,
		}

		transactionId, err := usecase.Do(*input)

		if err != nil {
			panic("Failed to execute use case")
		}

		var response = &NewAssetResponse{
			TransactionId: transactionId,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusCreated, string(jsonResponse))
	})
}

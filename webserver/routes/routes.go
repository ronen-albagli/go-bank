package router

import (
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
}

func CreateServerRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, "{}")
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

		usecase := &usecase.AddCreditUseCaseStruct{
			&configuration,
		}

		err := usecase.Do(*input)

		if err != nil {
			panic("Failed to execute use case")
		}

		c.JSON(http.StatusCreated, parsedInput)
	})
}

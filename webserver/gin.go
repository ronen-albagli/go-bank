package webserver

import (

	// "net/http"

	consumers "bank/webserver/consumers"
	router "bank/webserver/routes"

	"github.com/gin-gonic/gin"
)

// func CreateServerRoutes(router *gin.Engine) {
// 	router.GET("/", func(c *gin.Context) {
// 		fmt.Println("sdfdsf")
// 		c.JSON(http.StatusCreated, "{}")
// 	})
// }

func CreateWebServer() *gin.Engine {
	app := gin.Default()

	router.CreateServerRoutes(app)
	consumers.InvoiceConsumer()

	return app
}

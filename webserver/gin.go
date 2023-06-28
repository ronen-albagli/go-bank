package webserver

import (

	// "net/http"

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

	// go consumers.InvoiceConsumer()

	return app
}

package webserver

import (
	"bank/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServerRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, "{}")
	})
}

func CreateWebServer() *gin.Engine {
	app := gin.Default()

	conf := new(config.Config)

	conf.SetConf()

	CreateServerRoutes(app)

	// mongo, err := conf.GetMongoClient()
	// if err != nil {
	// 	panic(err)
	// }

	return app
}

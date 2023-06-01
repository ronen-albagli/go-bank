package webserver

import (
	"github.com/gin-gonic/gin"
)

func CreateWebServer() *gin.Engine {
	app := gin.Default()

	return app
}

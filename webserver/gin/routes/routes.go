package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServerRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, "{}")
	})
}

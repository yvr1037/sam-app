package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func to define all http api
func InitHandelApi(g *gin.Engine) error {
	api := g.Group("/api")
	api.GET("/hello", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, map[string]interface{}{"data": "world"}) })

	client := api.Group("/client")
	{
		client.POST("/register", RegisterClient)
		client.POST("/change", ChangeClient)
		client.POST("/delete", DeleteClient)
	}
	adviser := api.Group("/adviser")
	{
		adviser.POST("/register")
	}
	return nil
}

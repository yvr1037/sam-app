package handle

import (
	"hello-world/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func to define all http api
func InitHandelApi(g *gin.Engine) error {
	api := g.Group("/api")
	api.POST("/hello", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, map[string]interface{}{"data": "world"}) })
	client := api.Group("/client")
	{
		client.POST("/register", RegisterClient)
		client.POST("/signin", SignInClient)
		client.POST("/change", util.Auth(), ChangeClient)
		client.POST("/adviser/list", util.Auth(), GetUserAdviserList)
		client.POST("/adviser/info", util.Auth(), GetUserAdviserInfo)
	}
	adviser := api.Group("/adviser")
	{
		adviser.POST("/register", RegisterAdviser)
		adviser.POST("/signin", SignInAdviser)
		adviser.POST("/change", util.Auth(), ChangeAdvister)
	}
	order := api.Group("/order")
	{
		order.POST("/create", util.Auth())
		order.POST("/info", util.Auth())
		order.POST("/change/status", util.Auth())
	}
	return nil
}

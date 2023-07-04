package handle

import (
	"hello-world/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeClient(c *gin.Context) {
	cli := &model.Client{}
	err := c.BindJSON(cli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	err = cli.Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "ok"})
}

func RegisterClient(c *gin.Context) {
	cli := &model.Client{}
	err := c.BindJSON(cli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	err = cli.Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "ok"})
}

func DeleteClient(c *gin.Context) {
	type Name struct {
		Name string `json:"name"`
	}
	var name Name
	err := c.BindJSON(&name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	err = model.DeleteClient(name.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "ok"})
}

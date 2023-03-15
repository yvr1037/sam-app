package handle

import (
	"hello-world/internal/model"
	"hello-world/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ClientRole = 0

func ChangeClient(c *gin.Context) {
	cli := &model.Client{}
	err := c.BindJSON(cli)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err})
		return
	}

	//check if user exist
	exist, err := model.ExistClient(cli.Phone)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		Error(c, http.StatusBadRequest, "user not exist")
		return
	}

	cli.Password = util.SHA1(cli.Password + util.Salt)
	err = cli.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": "ok"})
}

func RegisterClient(c *gin.Context) {
	cli := &model.Client{}
	err := c.BindJSON(cli)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	//check if user exist
	exist, err := model.ExistClient(cli.Phone)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if exist {
		Error(c, http.StatusBadRequest, "user exist")
		return
	}

	cli.Password = util.SHA1(cli.Password + util.Salt)
	err = cli.Insert()
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	jwt, err := util.CreateJwt(cli.Phone, ClientRole)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
	}
	Success(c, map[string]interface{}{"data": jwt})
}

func SignInClient(c *gin.Context) {
	type SignIn struct {
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}
	var info SignIn
	err := c.BindJSON(&info)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	info.Password = util.SHA1(info.Password + util.Salt)
	exist, err := model.LoginByPwd(info.Password, info.Phone)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		Error(c, http.StatusBadRequest, "no such user")
		return
	}

	jwt, err := util.CreateJwt(info.Phone, ClientRole)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
	}
	Success(c, map[string]interface{}{"data": jwt})
}

func GetUserAdviserList(g *gin.Context) {
	temp, exist := g.Get("uid")
	uid := temp.(string)
	if !exist || len(uid) == 0 {
		Error(g, http.StatusBadRequest, "invalue uid")
		return
	}

	client, err := model.InfoClient(uid)
	if err != nil {
		Error(g, http.StatusBadRequest, err.Error())
		return
	}

	Success(g, map[string]interface{}{
		"list": client.Advisers,
	})
}

func GetUserAdviserInfo(g *gin.Context) {
	type Phone struct {
		Phone string `json:"phone"`
	}
	var info Phone
	err := g.BindJSON(&info)
	if err != nil {
		Error(g, http.StatusBadRequest, err.Error())
		return
	}

	data, err := model.InfoAdviser(info.Phone)
	if err != nil {
		Error(g, http.StatusBadRequest, err.Error())
		return
	}

	//deal secret message
	data.Password = ""

	Success(g, map[string]interface{}{
		"info": data,
	})
}

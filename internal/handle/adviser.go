package handle

import (
	"hello-world/internal/model"
	"hello-world/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const AdviserRole = 1

func ChangeAdvister(c *gin.Context) {
	adviser := &model.Adviser{}
	err := c.BindJSON(adviser)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err})
		return
	}

	//check if user exist
	exist, err := model.ExistAdvister(adviser.Phone)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		Error(c, http.StatusBadRequest, "user not exist")
		return
	}

	adviser.Password = util.SHA1(adviser.Password + util.Salt)
	err = adviser.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": "ok"})
}

func RegisterAdviser(c *gin.Context) {
	adviser := &model.Adviser{}
	err := c.BindJSON(adviser)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	//check if user exist
	exist, err := model.ExistAdvister(adviser.Phone)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if exist {
		Error(c, http.StatusBadRequest, "adviser exist")
		return
	}

	adviser.Password = util.SHA1(adviser.Password + util.Salt)
	err = adviser.Insert()
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	jwt, err := util.CreateJwt(adviser.Phone, AdviserRole)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
	}
	Success(c, map[string]interface{}{"data": jwt})
}

func SignInAdviser(c *gin.Context) {
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

	jwt, err := util.CreateJwt(info.Phone, AdviserRole)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
	}
	Success(c, map[string]interface{}{"data": jwt})
}

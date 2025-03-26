package apiControl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"github.com/gin-gonic/gin"
)

type jscode2sessionParams struct {
	Code   string
	Appid  string
	Secret string
}

func Dylogin(c *gin.Context) {
	var obj = bindtypes.Dylogin{}
	c.ShouldBind(&obj)

	//token, _ := service.GenerateToken(obj.Username, obj.Password)
	jscode2sessionParams1 := jscode2sessionParams{
		Code:   c.Query("code"),
		Appid:  "tte47d92c07b4bdaf701",
		Secret: "04811e0f85f1c78d4957c12591aab70ebc933130",
	}
	jscode2sessionParams1Json, _ := json.Marshal(jscode2sessionParams1)
	payload := strings.NewReader(string(jscode2sessionParams1Json))
	response, err := http.Post("https://developer.toutiao.com/api/apps/v2/jscode2session", "application/json", payload)

	if err != nil {
		c.JSON(200, gin.H{
			"code": "400",
			"msg":  "错误",
		})
		return
	}

	bytes, err1 := io.ReadAll(response.Body)
	if err1 != nil {
		c.JSON(200, gin.H{
			"code": "400",
			"msg":  "错误",
		})
		return
	}
	fmt.Println("Response:", string(bytes))
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
		"data": string(bytes),
	})
	defer response.Body.Close()

}

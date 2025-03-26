package apiControl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type clientTokenParams struct {
	Client_key    string
	Client_secret string
	Grant_type    string
}

func ClientToken(c *gin.Context) {
	var obj = clientTokenParams{}
	c.ShouldBind(&obj)

	// token, _ := service.GenerateToken(obj.Username, obj.Password)
	jscode2sessionParams1 := clientTokenParams{
		Client_key:    "tte47d92c07b4bdaf701",
		Client_secret: "04811e0f85f1c78d4957c12591aab70ebc933130",
		Grant_type:    "client_credential",
	}
	jscode2sessionParams1Json, _ := json.Marshal(jscode2sessionParams1)
	payload := strings.NewReader(string(jscode2sessionParams1Json))
	response, err := http.Post("https://open.douyin.com/oauth/client_token/", "application/json", payload)

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

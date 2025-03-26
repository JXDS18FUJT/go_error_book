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

type schemav2Params struct {
	AccessToken string
}

func Schemav2(c *gin.Context) {
	var obj = bindtypes.Dylogin{}
	c.ShouldBind(&obj)

	//token, _ := service.GenerateToken(obj.Username, obj.Password)
	schemav2Params1 := schemav2Params{
		AccessToken: "clt.f2f1437940f0342c416762bfda452f5aQHx1takmQWo4oWwSPQeo6XW7K5eN_lq",
	}
	schemav2Params1Json, _ := json.Marshal(schemav2Params1)
	payload := strings.NewReader(string(schemav2Params1Json))
	response, err := http.Post("https://open.douyin.com/api/apps/v1/url/generate_schema/", "application/json", payload)

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

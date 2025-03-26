package apiValid

import (
	"regexp"

	"347613781qq.com/goInit1/app/http/control/apiControl"
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
)

func UploadValid(c *gin.Context) {
	file, fileErr := c.FormFile("file")
	name, _ := c.GetPostForm("name")
	saveErr := c.SaveUploadedFile(file, "storage/app/img/"+name+file.Filename)
	if fileErr != nil {

		utils.ResError(402, fileErr.Error(), gin.H{}, c)
		return

	}
	if saveErr != nil {

		utils.ResError(402, saveErr.Error(), gin.H{}, c)
		return
	}
	// matchRes, _ := regexp.MatchString(`(.png|.jpeg|.jpg)$`, file.Filename)
	// fmt.Printf("%s,是否通过检验:%v \n", file.Filename, matchRes)
	//file.Header["Content-Type"]
	matchRes, _ := regexp.MatchString(`(.png|.jpeg|.jpg)$`, file.Filename)
	if !matchRes {
		utils.ResError(402, "文件类型错误", nil, c)
		return
	}
	apiControl.Upload(c)

}

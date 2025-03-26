package apiControl

import (
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 上传图片的接口
// @Description get string by ID
// @Accept png,jpeg
// @Produce  json
// @Param  file formData file true  "文件"
// @Param name formData  string false "名字"

// @Success 200 {object} model.Nong	"ok"
// @Failure 400 {object} global.AppBusinessError
// @Router /upload [post]
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	name, _ := c.GetPostForm("name")
	saveErr := c.SaveUploadedFile(file, "storage/app/img/"+name+file.Filename)
	if saveErr != nil {

		utils.ResError(402, saveErr.Error(), gin.H{}, c)
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
		"name": file.Filename,
	})
}

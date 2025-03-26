package apiControl

import (
	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
)

// @Summary nong的接口
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param id query int true "Id"
// @Success 200 {object} model.Nong	"ok"
// @Failure 400 {object} global.AppBusinessError
// @Router /nong [get]
func Nong(c *gin.Context) {
	query := bindtypes.Nongs{}
	c.ShouldBindUri(&query)

	var nongs, nongsErr = model.GetNong(query.Id)

	if nongsErr == nil {
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "成功",
			"data": nongs,
		})
	} else {
		c.JSON(200, gin.H{
			"code": "400",
			"msg":  "" + nongsErr.Error(),
		})
	}

}

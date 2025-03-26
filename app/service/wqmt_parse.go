package service

import (
	"strings"

	"347613781qq.com/goInit1/app/model"
)

func WqmtParseUpPool(datas []string) {
	if len(datas) != 7 {
		return
	}

	character := datas[1]
	remark := datas[3]
	upPool := model.UpPool{}
	upPool.Name = character
	upPool.Remark = remark
	if strings.Contains(remark, "常规追踪") {
		upPool.Kind = 1
		upPool.KindName = "常规追踪"

	} else if strings.Contains(remark, "活动追踪") {
		upPool.Kind = 2
		upPool.KindName = "活动追踪"

	} else if strings.Contains(remark, "限定追踪") {
		upPool.Kind = 3
		upPool.KindName = "限定追踪"

	} else if strings.Contains(remark, "统合追踪") {
		upPool.Kind = 4
		upPool.KindName = "统合追踪"

	}

	model.CreateUpPool(&upPool)

}

func WqmtParseCharacterSkill(datas []string) {

}

package model

import (
	"347613781qq.com/goInit1/utils"
)

type UpPool struct {
	Id                  int `gorm:"AUTO_INCREMENT;DEFAULT:0"`
	Name                string
	Remark              string
	Kind                int
	KindName            string
	PoolCharacterImgUrl string
}

func UpdateUpPool(upPool *UpPool) (err error) {
	err = utils.DB.Save(upPool).Error
	return
}
func CreateUpPool(upPool *UpPool) (err error) {
	err = utils.DB.Create(upPool).Error
	return
}
func DeleteAllUpPool(upPool UpPool) (err error) {
	utils.DB.Where("id > 0").Delete(upPool)

	return
}
func TruncateUpPool() (err error) {
	utils.DB.Exec("TRUNCATE `up_pools`")

	return
}

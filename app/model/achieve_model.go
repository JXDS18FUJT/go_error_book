package model

import "347613781qq.com/goInit1/utils"

type Achieve struct {
	Id   int
	Name string

	ImgUrl     string
	Level      int
	LevelName  string
	UnlockDesc string

	Access string
}

func UpdateAchieve(Achieve *Achieve) (err error) {
	err = utils.DB.Save(Achieve).Error
	return
}
func CreateAchieve(Achieve *Achieve) (err error) {
	err = utils.DB.Create(Achieve).Error
	return
}
func DeleteAllAchieve(Achieve Achieve) (err error) {
	utils.DB.Where("id > 0").Delete(Achieve)

	return
}
func TruncateAchieve() (err error) {
	utils.DB.Exec("TRUNCATE `achieves`")

	return
}

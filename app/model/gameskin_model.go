package model

import "347613781qq.com/goInit1/utils"

type Gameskin struct {
	Id         int
	Name       string
	Character  string
	ImgUrl     string
	Level      int
	LevelName  string
	Series     string
	OnlineTime string
	Access     string
}

func UpdateGameskin(Gameskin *Gameskin) (err error) {
	err = utils.DB.Save(Gameskin).Error
	return
}
func CreateGameskin(Gameskin *Gameskin) (err error) {
	err = utils.DB.Create(Gameskin).Error
	return
}
func DeleteAllGameskin(Gameskin Gameskin) (err error) {
	utils.DB.Where("id > 0").Delete(Gameskin)

	return
}
func TruncateGameskin() (err error) {
	utils.DB.Exec("TRUNCATE `gameskins`")

	return
}

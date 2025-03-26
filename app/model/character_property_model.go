package model

import (
	"347613781qq.com/goInit1/utils"
)

//设置主键

type CharacterProperty struct {
	CharacterPropertyId int    `json:"character_property_id"`
	Character           string `json:"character"`
	Code                string `json:"code"`
	Career              string `json:"career"`
	Camp                string `json:"camp"`
	Hp                  string `json:"hp"`
	Atk                 string `json:"atk"`
	Def                 string `json:"def"`
	Rgs                 string `json:"rgs"`
	Block               string `json:"block"`
	Speed               string `json:"speed"`
	Tag                 string `json:"tag"`
	LevelName           string `json:"level_name"`
	Level               int    `json:"level"`
	HeadImage           string `json:"head_image"`
}

func GetCharacterPropertyFirst(character *CharacterProperty) (firstCharacter CharacterProperty, err error) {
	firstCharacter = CharacterProperty{}
	utils.DB.Where(character).First(firstCharacter)
	return firstCharacter, err

}

func CreateCharacterProperty(character *CharacterProperty) (err error) {
	err = utils.DB.Create(&character).Error
	return
}
func UpdateCharacterProperty(character *CharacterProperty) (err error) {
	err = utils.DB.Save(&character).Error
	return
}
func TruncateCharacterProperty() (err error) {
	utils.DB.Exec("TRUNCATE `character_properties`")

	return
}

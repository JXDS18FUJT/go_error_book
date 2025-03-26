package model

import "347613781qq.com/goInit1/utils"

type CharacterSkill struct {
	Id                int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;DEFAULT:0"`
	Character         string
	Code              string
	BasicAttackName   string
	BasicAttackDesc   string
	BasicAttackLvDesc string

	ChargeSkillName   string
	ChargeSkillTag    string
	ChargeSkillDesc   string
	ChargeSkillLvDesc string

	PassiveSkill1Name   string
	PassiveSkill1Desc   string
	PassiveSkill1LvDesc string

	PassiveSkill2Name   string
	PassiveSkill2Desc   string
	PassiveSkill2LvDesc string
}

func FirstOrCreateCharacterSkill(CharacterSkill *CharacterSkill) (err error) {
	err = utils.DB.FirstOrCreate(CharacterSkill).Error
	return
}
func SaveCharacterSkill(CharacterSkill *CharacterSkill) (err error) {
	err = utils.DB.Save(CharacterSkill).Error
	return
}
func UpdatesCharacterSkill(CharacterSkill *CharacterSkill) (err error) {
	err = utils.DB.Model(CharacterSkill).Updates(CharacterSkill).Error
	return
}
func CreateCharacterSkill(CharacterSkill *CharacterSkill) (err error) {
	err = utils.DB.Create(CharacterSkill).Error
	return
}
func DeleteAllCharacterSkill(CharacterSkill CharacterSkill) (err error) {
	utils.DB.Where("id > 0").Delete(CharacterSkill)

	return
}
func TruncateCharacterSkill() (err error) {
	utils.DB.Exec("TRUNCATE `CharacterSkills`")

	return
}

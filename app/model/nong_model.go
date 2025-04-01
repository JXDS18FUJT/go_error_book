package model

import (
	"347613781qq.com/goInit1/app/global"
	"347613781qq.com/goInit1/utils"
)

//设置主键

type Nong struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Content   string       `json:"content"`
	CreatedAt global.XTime `json:"created_at"`
	UpdatedAt global.XTime `json:"updated_at"`
}

func CreateNong(nong *Nong) (err error) {
	err = utils.DB.Create(&nong).Error
	return
}
func UpdateNong(nong *Nong) (err error) {
	err = utils.DB.Save(&nong).Error
	return
}
func DeleteNong(id int) (err error) {

	err = utils.DB.First("id=?", id).Delete(&Nong{}).Error
	return
}
func GetNong(id int) (nong Nong, err error) {
	if utils.DB == nil {
		return Nong{}, nil
	}

	err = utils.DB.Limit(1).Where("id = ?", id).First(&nong).Error
	return nong, err

}
func GetAllNong() (nongList []Nong, err error) {

	if err = utils.DB.Find(&nongList).Error; err != nil {
		return nil, err
	}
	return
}

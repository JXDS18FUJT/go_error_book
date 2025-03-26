package model

import "347613781qq.com/goInit1/utils"

type ExBrand struct {
	Id int `json:"id"`

	Name      string `json:"name"`
	Character string `json:"character"`
	ImgUrl    string `json:"img_url"`
	Effect    string `json:"effect"`
}

func UpdateExBrand(exBrand *ExBrand) (err error) {
	err = utils.DB.Save(exBrand).Error
	return
}
func CreateExBrand(exBrand *ExBrand) (err error) {
	err = utils.DB.Create(exBrand).Error
	return
}
func DeleteAllExBrand(exBrand ExBrand) (err error) {
	utils.DB.Where("id > 0").Delete(exBrand)

	return
}
func TruncateExBrand() (err error) {
	utils.DB.Exec("TRUNCATE `ex_brands`")

	return
}

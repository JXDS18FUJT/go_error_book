package model

import "347613781qq.com/goInit1/utils"

type ComBrand struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ImgUrl    string `json:"img_url"`
	Suit1     string `json:"suit1"`
	Suit2     string `json:"suit2"`
	Suit3     string `json:"suit3"`
	Effect    string `json:"effect"`
	Desc      string `json:"desc"`
	LevelName string `json:"level_name"`
	Level     int    `json:"level"`
	Access    string `json:"access"`
}

func UpdateComBrand(comBrand *ComBrand) (err error) {
	err = utils.DB.Save(comBrand).Error
	return
}
func CreateComBrand(comBrand *ComBrand) (err error) {
	err = utils.DB.Create(comBrand).Error
	return
}
func DeleteAllComBrand(comBrand ComBrand) (err error) {
	utils.DB.Where("id > 0").Delete(comBrand)

	return
}
func TruncateComBrand() (err error) {
	utils.DB.Exec("TRUNCATE `com_brands`")

	return
}

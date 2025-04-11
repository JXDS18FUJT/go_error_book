package model

import (
	"347613781qq.com/goInit1/utils"
)

//设置主键

type Question struct {
	Uuid     int16  `form:"uuid" json:"uuid" validate:"required"`
	Id       string `form:"id" json:"id" validate:"required"`
	Question string `form:"question" json:"question" validate:"required"`
	Option1  string `form:"option1" json:"option1" validate:"required"`
	Option2  string `form:"option2" json:"option2" validate:"required"`
	Option3  string `form:"option3" json:"option3" validate:"required"`
	Option4  string `form:"option4" json:"option4" validate:"required"`
	Answer   string `form:"answer" json:"answer" validate:"required"`
	Image    string `form:"image" json:"image" validate:"required"`
}

func CreateQuestion(question *Question) (err error) {
	err = utils.DB.Create(&question).Error
	return
}
func CreateQuestions(questions *[]Question) (err error) {
	err = utils.DB.Create(questions).Error
	return
}
func UpdateQuestion(question *Question) (err error) {
	err = utils.DB.Save(&question).Error
	return
}
func DeleteQuestion(id int) (err error) {

	err = utils.DB.First("id=?", id).Delete(&Question{}).Error
	return
}
func GetQuestion(id int) (question Question, err error) {
	if utils.DB == nil {
		return Question{}, nil
	}

	err = utils.DB.Limit(1).Where("id = ?", id).First(&question).Error
	return question, err

}
func GetAllQuestion() (questionList []Question, err error) {

	if err = utils.DB.Find(&questionList).Error; err != nil {
		return nil, err
	}
	return
}

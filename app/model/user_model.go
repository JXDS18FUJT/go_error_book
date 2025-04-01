package model

import (
	"347613781qq.com/goInit1/app/global"
	"347613781qq.com/goInit1/utils"
)

type User struct {
	Id           int          `json:"id"`
	Username     string       `json:"username"`
	PasswordHash string       `json:"-"`
	Status       int          `json:"status"`
	CreatedAt    global.XTime `json:"created_at"`
	UpdatedAt    global.XTime `json:"updated_at"`
}

func CreateUser(user *User) (err error) {
	err = utils.DB.Create(&user).Error
	return
}
func UpdateUser(user *User) (err error) {
	err = utils.DB.Save(&user).Error
	return
}
func DeleteUser(id int) (err error) {

	err = utils.DB.First("id=?", id).Delete(&User{}).Error
	return
}
func GetUser(username string) (user User, err error) {
	if utils.DB == nil {
		return User{}, nil
	}

	err = utils.DB.Limit(1).Where("username = ?", username).First(&user).Error
	return user, err

}

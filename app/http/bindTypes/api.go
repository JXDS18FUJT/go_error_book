package bindtypes

type Login struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
type Nongs struct {
	Id int `json:"id" uri:"id"`
}

type Dylogin struct {
	Code string `form:"code" json:"code" validate:"required"`
}

type Question struct {
	Uuid     int16  `form:"uuid" json:"uuid" validate:"required"`
	Id       string `form:"id" json:"id" validate:"required"`
	Question string
	Option1  string
	Option2  string
	Option3  string
	Option4  string
	Answer   string
	Image    string
}

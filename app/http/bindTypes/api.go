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

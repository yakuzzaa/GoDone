package apiGateway

type User struct {
	Id       string `json:"-"`
	Name     string `json:"name"`
	Username string `json:"email"`
	Password string `json:"password"`
}

package form

type LoginForm struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

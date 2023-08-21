package model

type LoginInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type RegisterInput struct {
	Account         string `json:"account"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}
type GetCaptchaOut struct {
	Img string `json:"img"`
}

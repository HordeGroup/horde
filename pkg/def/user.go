package def

type UserRegisterRequest struct {
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	Telephone string `json:"telephone" form:"telephone"`
}

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

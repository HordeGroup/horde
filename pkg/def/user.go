package def

type UserRegisterRequest struct {
	Name      string `json:"user_name" form:"user_name" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Email     string `json:"email" form:"email"`
	Telephone string `json:"telephone" form:"telephone"`
}

type UserRegisterData struct {
	UserId uint32 `json:"user_id"`
}

type UserLoginRequest struct {
	Name     string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginData struct {
	UserId uint32 `json:"user_id"`
}

type NeedLoginData struct {
	NeedLogin bool `json:"need_login"`
}

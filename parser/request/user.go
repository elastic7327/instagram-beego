package request

type RegisterFormRequest struct {
	DisplayName string `form:"displayName"`
	Email       string `form:"email"`
	Passsword   string `form:"password"`
}

type LoginFormRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UpdateFormRequest struct {
	DisplayName string `form:"displayName"`
	Email       string `form:"email"`
}

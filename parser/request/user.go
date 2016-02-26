package request

type RegisterFormRequest struct {
	DisplayName string `form:"displayName"`
	Email       string `form:"email"`
	Passsword   string `form:"password"`
}

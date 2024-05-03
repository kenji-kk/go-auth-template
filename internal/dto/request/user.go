package request

type CreateUser struct {
	UserName string `json:"user_name"  binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=5,max=30"`
	Email    string `json:"email" binding:"required,min=5,max=100"`
}

type Login struct {
	Password string `json:"password"  binding:"required,min=5,max=30"`
	Email    string `json:"email" binding:"required,min=5,max=100"`
}

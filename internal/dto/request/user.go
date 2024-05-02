package request

type CreateUser struct {
	UserName string `json:"user_name" db:"user_name" binding:"required,min=5,max=30"`
	Password string `json:"password" db:"password" binding:"required,min=5,max=30"`
	Email    string `json:"email" db:"email" binding:"required,min=5,max=100"`
}
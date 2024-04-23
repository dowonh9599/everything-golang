package requests

type UserRequest struct {
	Id       int    `json:"id" form:"id" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

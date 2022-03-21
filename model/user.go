package model

type User struct {
	Id       int    `json:"id" form:"id" binding:"id"`
	UserName string `json:"username" form:"username" binding:"username"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
}

type UserRegister struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile" form:"mobile" binding:"mobile"`
	Email    string `json:"email"`
}

type LoginRequestParams struct {
	UserName string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	//Ctx      *gin.Context
}

type SelectUserByIdParam struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required"`
}

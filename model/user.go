package model

type UserRegister struct {
	Nama_User     string `json:"nama_user" binding:"required"`
	Email_User    string `json:"email_user" binding:"required,email"`
	Password_User string `json:"password_user" binding:"required,min=8"`
}

type UserLogin struct {
	Email_User    string `json:"email_user" binding:"required,email"`
	Password_User string `json:"password_user" binding:"required"`
}

type UserLoginResponses struct {
	Token  string `json:"token"`
	RoleID int    `json:"role_id"`
}

type UserParam struct {
	ID_User       int    `json:"-"`
	Email_User    string `json:"-"`
	Password_User string `json:"-"`
	RoleID        int    `json:"-"`
}

package model

type AdminRegister struct {
	Email_Admin    string `json:"email_admin" binding:"required,email"`
	Password_Admin string `json:"password_admin" binding:"required,min=8"`
}

type AdminLogin struct {
	Email_Admin    string `json:"email_admin" binding:"required,email"`
	Password_Admin string `json:"password_admin" binding:"required"`
}

type AdminLoginResponses struct {
	Token string `json:"token"`
}

type AdminParam struct {
	ID_Admin       int    `json:"-"`
	Email_Admin    string `json:"-"`
	Password_Admin string `json:"-"`
}

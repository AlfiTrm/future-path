package entity

type User struct {
	ID_User       int    `json:"id_user" gorm:"type:smallint;primary_key;autoIncrement"`
	Nama_User     string `json:"nama_user" gorm:"type:varchar(50);not null"`
	Email_User    string `json:"email_user" gorm:"type:varchar(30);not null"`
	Password_User string `json:"password_user" gorm:"type:varchar(100);not null"`
	Asal_Sekolah  string `json:"asal_sekolah" gorm:"type:varchar(30)"`
	RoleID        int    `json:"role_id"`
	Token         string `json:"token" gorm:"type:text"`
}

// esTeh123

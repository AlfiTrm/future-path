package entity

type Role struct {
	Id_role   int    `json:"id_role" gorm:"type:smallint;primary_key;autoIncrement"`
	Role_Name string `json:"role_name" gorm:"type:varchar(10);not null"`
	Users     []User `json:"users" gorm:"foreignKey:RoleID"`
}

package entity

type Universitas struct {
	ID_Universitas        int    `json:"id_universitas" gorm:"type:smallint;primary_key;autoIncrement"`
	Nama_Universitas      string `json:"nama_universitas" gorm:"type:varchar(50);not null"`
	Alamat_Universitas    string `json:"alamat_universitas" gorm:"type:varchar(30);not null"`
	Deskripsi_Universitas string `json:"deskripsi_universitas" gorm:"type:text;not null"`
	ID_Kepemilikan        int    `json:"id_kepemilikan"`
}

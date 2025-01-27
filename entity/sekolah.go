package entity

type Sekolah struct {
	ID_Sekolah        int    `json:"id_sekolah" gorm:"type:smallint;primary_key;autoIncrement"`
	Nama_Sekolah      string `json:"nama_sekolah" gorm:"type:varchar(50);not null"`
	Alamat_Sekolah    string `json:"alamat_sekolah" gorm:"type:varchar(30);not null"`
	Deskripsi_Sekolah string `json:"deskripsi_sekolah" gorm:"type:text;not null"`
	ID_Kepemilikan    int    `json:"id_kepemilikan"`
	PhotoLink         string `json:"photoLink" gorm:"type:varchar(200)"`
}

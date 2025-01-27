package entity

import "time"

type Berita struct {
	ID_Berita    int       `json:"id_berita" gorm:"type:smallint;primary_key;autoIncrement"`
	Judul_Berita string    `json:"judul_berita" gorm:"type:varchar(100);not null"`
	Isi_Berita   string    `json:"isi_berita" gorm:"type:text;not null"`
	Tanggal      time.Time `json:"tanggal" gorm:"autoCreateTime"`
}

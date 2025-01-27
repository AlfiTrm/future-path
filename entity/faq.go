package entity

type FAQ struct {
	ID_FAQ    int    `json:"id_faq" gorm:"type:smallint;primary_key;autoIncrement"`
	Judul_FAQ string `json:"judul_faq" gorm:"type:varchar(100);not null"`
	Isi_FAQ   string `json:"isi_faq" gorm:"type:text;not null"`
}

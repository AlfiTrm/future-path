package entity

type Kepemilikan struct {
	ID_Kepemilikan   int           `json:"id_kepemilikan" gorm:"type:smallint;primary_key;autoIncrement"`
	Nama_Kepemilikan string        `json:"nama_kepemilikan" gorm:"type:varchar(10);not null"`
	Sekolah          []Sekolah     `json:"sekolahs" gorm:"foreignKey:ID_Kepemilikan"`
	Universitas      []Universitas `json:"universitas" gorm:"foreignKey:ID_Kepemilikan"`
}

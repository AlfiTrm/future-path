package model

type CreateBerita struct {
	Judul_Berita string `json:"judul_berita" binding:"required"`
	Isi_Berita   string `json:"isi_berita" binding:"required"`
}

type GetBerita struct {
	ID_Berita    int    `json:"id_berita"`
	Judul_Berita string `json:"judul_berita"`
	Isi_Berita   string `json:"isi_berita"`
}

type UpdateBerita struct {
	Judul_Berita string `json:"judul_berita"`
	Isi_Berita   string `json:"isi_berita"`
}

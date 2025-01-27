package model

type GetUniversitas struct {
	Nama_Universitas      string `json:"nama_universitas"`
	Alamat_Universitas    string `json:"alamat_universitas"`
	Deskripsi_Universitas string `json:"deskripsi_universitas"`
}

type GetAllUniv struct {
	Nama_Universitas   string `json:"nama_universitas"`
	Alamat_Universitas string `json:"alamat_universitas"`
}

type CreateUniv struct {
	Nama_Universitas      string `json:"nama_universitas" binding:"required"`
	Alamat_Universitas    string `json:"alamat_universitas" binding:"required"`
	Deskripsi_Universitas string `json:"deskripsi_universitas" binding:"required"`
	ID_Kepemilikan        int    `json:"id_kepemilikan" binding:"required"`
}

type CreateUnivResponse struct {
	Nama_Universitas      string `json:"nama_universitas"`
	Alamat_Universitas    string `json:"alamat_universitas"`
	Deskripsi_Universitas string `json:"deskripsi_universitas"`
	ID_Kepemilikan        int    `json:"id_kepemilikan"`
}

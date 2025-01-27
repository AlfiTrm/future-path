package service

import (
	"errors"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
	"future-path/pkg/supabase"
)

type ISekolahService interface {
	GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error)
	GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error)
	GetAllSekolah(page int) ([]*entity.Sekolah, error, int64)
	GetSekolahDetail(id int) (*entity.Sekolah, error)
	AddSekolah(sekolahReq *model.CreateSekolah, param model.UploadPhoto) (*entity.Sekolah, error)
}

type SekolahService struct {
	SekolahRepository repository.ISekolahRepository
	supabase          supabase.Interface
}

func NewSekolahService(SekolahRepository repository.ISekolahRepository, supabase supabase.Interface) ISekolahService {
	return &SekolahService{
		SekolahRepository: SekolahRepository,
		supabase:          supabase,
	}
}

func (ss *SekolahService) GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahNegeri(namaSekolah)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (ss *SekolahService) GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahSwasta(namaSekolah)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (ss *SekolahService) GetAllSekolah(page int) ([]*entity.Sekolah, error, int64) {
	limit := 10
	offset := (page - 1) * limit
	sekolah, err := ss.SekolahRepository.GetAllSekolah(limit, offset)
	if err != nil {
		return nil, err, 0
	}

	totalData, err := ss.SekolahRepository.CountAllSekolah()
	if err != nil {
		return nil, err, 0
	}

	return sekolah, nil, totalData
}

func (ss *SekolahService) GetSekolahDetail(id int) (*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahDetail(id)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (ss *SekolahService) AddSekolah(sekolahReq *model.CreateSekolah, param model.UploadPhoto) (*entity.Sekolah, error) {

	link, err := ss.supabase.Upload(param.Photo)
	if err != nil {
		return nil, err
	}

	sekolah := &entity.Sekolah{
		Nama_Sekolah:      sekolahReq.Nama_Sekolah,
		Alamat_Sekolah:    sekolahReq.Alamat_Sekolah,
		Deskripsi_Sekolah: sekolahReq.Deskripsi_Sekolah,
		ID_Kepemilikan:    sekolahReq.ID_Kepemilikan,
		PhotoLink:         link,
	}

	if sekolah.ID_Kepemilikan != 1 && sekolah.ID_Kepemilikan != 2 {
		return nil, errors.New("invalid ownerships id")
	}

	sekolah, err = ss.SekolahRepository.AddSekolah(sekolah)
	if err != nil {
		_ = ss.supabase.Delete(link)
		return nil, err
	}
	return sekolah, nil
}

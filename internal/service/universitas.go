package service

import (
	"errors"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
)

type IUniversitasService interface {
	GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error)
	GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error)
	GetAllUniv(page int) ([]*entity.Universitas, error)
	GetUnivDetail(id int) (*entity.Universitas, error)
	AddUniv(univReq *model.CreateUniv) (*entity.Universitas, error)
}

type UniversitasService struct {
	UniversitasRepository repository.IUniversitasRepository
}

func NewUniversitasService(UniversitasRepository repository.IUniversitasRepository) IUniversitasService {
	return &UniversitasService{UniversitasRepository}
}

func (us *UniversitasService) GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error) {
	univ, err := us.UniversitasRepository.GetUnivNegeri(namaUniv)
	if err != nil {
		return nil, err
	}
	return univ, nil
}

func (us *UniversitasService) GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error) {
	univ, err := us.UniversitasRepository.GetUnivSwasta(namaUniv)
	if err != nil {
		return nil, err
	}
	return univ, nil
}

func (us *UniversitasService) GetAllUniv(page int) ([]*entity.Universitas, error) {
	limit := 10
	offset := (page - 1) * limit
	univ, err := us.UniversitasRepository.GetAllUniv(limit, offset)
	if err != nil {
		return nil, err
	}

	return univ, nil
}

func (us *UniversitasService) GetUnivDetail(id int) (*entity.Universitas, error) {
	univ, err := us.UniversitasRepository.GetUnivDetail(id)
	if err != nil {
		return nil, err
	}
	return univ, nil
}

func (us *UniversitasService) AddUniv(univReq *model.CreateUniv) (*entity.Universitas, error) {
	univ := &entity.Universitas{
		Nama_Universitas:      univReq.Nama_Universitas,
		Alamat_Universitas:    univReq.Alamat_Universitas,
		Deskripsi_Universitas: univReq.Deskripsi_Universitas,
		ID_Kepemilikan:        univReq.ID_Kepemilikan,
	}

	if univ.ID_Kepemilikan != 1 && univ.ID_Kepemilikan != 2 {
		return nil, errors.New("invalid ownerships id")
	}

	univ, err := us.UniversitasRepository.AddUniv(univ)
	if err != nil {
		return nil, err
	}
	return univ, nil
}

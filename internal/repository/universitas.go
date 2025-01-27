package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type IUniversitasRepository interface {
	GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error)
	GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error)
	GetAllUniv(limit, offset int) ([]*entity.Universitas, error)
	GetUnivDetail(id int) (*entity.Universitas, error)
	AddUniv(univ *entity.Universitas) (*entity.Universitas, error)
}

type UniversitasRepository struct {
	db *gorm.DB
}

func NewUniversitasRepository(db *gorm.DB) IUniversitasRepository {
	return &UniversitasRepository{db}
}

func (u *UniversitasRepository) GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error) {
	var univ []*entity.Universitas
	query := "%" + namaUniv + "%"
	if err := u.db.Debug().Where("id_kepemilikan = ? AND nama_universitas LIKE ?", 1, query).Find(&univ).Error; err != nil {
		return nil, err
	}
	return univ, nil
}

func (u *UniversitasRepository) GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error) {
	var univ []*entity.Universitas
	query := "%" + namaUniv + "%"
	if err := u.db.Debug().Where("id_kepemilikan = ? AND nama_universitas LIKE ?", 2, query).Find(&univ).Error; err != nil {
		return nil, err
	}
	return univ, nil
}

func (u *UniversitasRepository) GetAllUniv(limit, offset int) ([]*entity.Universitas, error) {
	var univ []*entity.Universitas
	if err := u.db.Debug().Limit(limit).Offset(offset).Find(&univ).Error; err != nil {
		return nil, err
	}
	return univ, nil
}

func (u *UniversitasRepository) GetUnivDetail(id int) (*entity.Universitas, error) {
	var univ entity.Universitas
	if err := u.db.Debug().Where("id_universitas = ?", id).First(&univ).Error; err != nil {
		return nil, err
	}
	return &univ, nil
}

func (u *UniversitasRepository) AddUniv(univ *entity.Universitas) (*entity.Universitas, error) {
	err := u.db.Debug().Create(&univ).Error
	if err != nil {
		return nil, err
	}
	return univ, nil
}

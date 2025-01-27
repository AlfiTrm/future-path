package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type ISekolahRepository interface {
	GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error)
	GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error)
	GetAllSekolah(limit, offset int) ([]*entity.Sekolah, error)
	GetSekolahDetail(id int) (*entity.Sekolah, error)
	AddSekolah(sekolah *entity.Sekolah) (*entity.Sekolah, error)
	CountAllSekolah() (int64, error)
}

type SekolahRepository struct {
	db *gorm.DB
}

func NewSekolahRepository(db *gorm.DB) ISekolahRepository {
	return &SekolahRepository{db}
}

func (sk *SekolahRepository) GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	query := "%" + namaSekolah + "%"
	if err := sk.db.Debug().Where("id_kepemilikan = ? AND nama_sekolah LIKE ?", 1, query).Find(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	query := "%" + namaSekolah + "%"
	if err := sk.db.Debug().Where("id_kepemilikan = ? AND nama_sekolah LIKE ?", 2, query).Find(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) GetAllSekolah(limit, offset int) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	if err := sk.db.Debug().Limit(limit).Offset(offset).Find(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) CountAllSekolah() (int64, error) {
	var count int64
	err := sk.db.Debug().Model(&entity.Sekolah{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (sk *SekolahRepository) GetSekolahDetail(id int) (*entity.Sekolah, error) {
	var sekolah entity.Sekolah
	if err := sk.db.Debug().Where("id_sekolah = ?", id).First(&sekolah).Error; err != nil {
		return nil, err
	}
	return &sekolah, nil
}

func (sk *SekolahRepository) AddSekolah(sekolah *entity.Sekolah) (*entity.Sekolah, error) {
	err := sk.db.Debug().Create(&sekolah).Error
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) UploadPhoto(photoLink string, id int) error {
	err := sk.db.Debug().Model(&entity.Sekolah{}).Where("id_sekolah = ?", id).Update("photoLink", photoLink).Error
	if err != nil {
		return err
	}
	return nil
}

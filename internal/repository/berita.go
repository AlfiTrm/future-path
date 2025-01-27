package repository

import (
	"future-path/entity"
	"future-path/model"

	"gorm.io/gorm"
)

type IBeritaRepository interface {
	CreateBerita(berita *entity.Berita) (*entity.Berita, error)
	GetBeritaSingkat(limit, offset int) ([]*entity.Berita, error)
	GetBeritaFull(id int) (*entity.Berita, error)
	UpdateBerita(id int, beritaRequest *model.UpdateBerita) (*entity.Berita, error)
	DeleteBerita(id int) error
	CountAllBerita() (int64, error)
}

type BeritaRepository struct {
	db *gorm.DB
}

func NewBeritaRepository(db *gorm.DB) IBeritaRepository {
	return &BeritaRepository{db}
}

func (br *BeritaRepository) CreateBerita(berita *entity.Berita) (*entity.Berita, error) {
	if err := br.db.Debug().Create(berita).Error; err != nil {
		return nil, err
	}

	return berita, nil
}

func (br *BeritaRepository) GetBeritaSingkat(limit, offset int) ([]*entity.Berita, error) {
	var berita []*entity.Berita

	if err := br.db.Debug().Select("id_berita, judul_berita, LEFT(isi_berita, 100) as isi_berita").Order("tanggal desc").Limit(limit).Offset(offset).Find(&berita).Error; err != nil {
		return nil, err
	}
	return berita, nil
}

func (br *BeritaRepository) CountAllBerita() (int64, error) {
	var count int64
	err := br.db.Debug().Model(&entity.Berita{}).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}

func (br *BeritaRepository) GetBeritaFull(id int) (*entity.Berita, error) {
	var berita entity.Berita
	if err := br.db.Debug().Where("id_berita = ?", id).First(&berita).Error; err != nil {
		return nil, err
	}

	return &berita, nil
}

func (br *BeritaRepository) UpdateBerita(id int, beritaRequest *model.UpdateBerita) (*entity.Berita, error) {
	tx := br.db.Begin()

	var berita entity.Berita
	if err := tx.Debug().Where("id_berita = ?", id).First(&berita).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	beritaParse := *parseUpdateBerita(beritaRequest, &berita)

	if err := tx.Debug().Save(&beritaParse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &berita, nil
}

func parseUpdateBerita(model *model.UpdateBerita, berita *entity.Berita) *entity.Berita {
	if model.Judul_Berita != "" {
		berita.Judul_Berita = model.Judul_Berita
	}

	if model.Isi_Berita != "" {
		berita.Isi_Berita = model.Isi_Berita
	}

	return berita
}

func (br *BeritaRepository) DeleteBerita(id int) error {
	tx := br.db.Begin()

	if err := tx.Debug().Where("id_berita = ?", id).First(&entity.Berita{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Where("id_berita = ?", id).Delete(&entity.Berita{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

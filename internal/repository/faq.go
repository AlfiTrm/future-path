package repository

import (
	"future-path/entity"
	"future-path/model"

	"gorm.io/gorm"
)

type IFAQRepository interface {
	CreateFAQ(faq *entity.FAQ) (*entity.FAQ, error)
	GetFAQ() ([]*entity.FAQ, error)
	UpdateFAQ(id int, FAQRequest *model.UpdateFAQ) (*entity.FAQ, error)
	DeleteFAQ(id int) error
}

type FAQRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) IFAQRepository {
	return &FAQRepository{db}
}

func (fr *FAQRepository) CreateFAQ(faq *entity.FAQ) (*entity.FAQ, error) {
	if err := fr.db.Debug().Create(faq).Error; err != nil {
		return nil, err
	}

	return faq, nil
}

func (fr *FAQRepository) GetFAQ() ([]*entity.FAQ, error) {
	var faq []*entity.FAQ

	if err := fr.db.Debug().Find(&faq).Error; err != nil {
		return nil, err
	}

	return faq, nil
}

func (fr *FAQRepository) UpdateFAQ(id int, FAQRequest *model.UpdateFAQ) (*entity.FAQ, error) {
	tx := fr.db.Begin()

	var faq entity.FAQ
	if err := tx.Debug().Where("id_faq = ?", id).First(&faq).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	faqParse := *parseUpdateFAQ(FAQRequest, &faq)

	if err := tx.Debug().Save(&faqParse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &faq, nil
}

func parseUpdateFAQ(model *model.UpdateFAQ, faq *entity.FAQ) *entity.FAQ {
	if model.Judul_FAQ != "" {
		faq.Judul_FAQ = model.Judul_FAQ
	}

	if model.Isi_FAQ != "" {
		faq.Isi_FAQ = model.Isi_FAQ
	}

	return faq
}

func (fr *FAQRepository) DeleteFAQ(id int) error {
	tx := fr.db.Begin()

	if err := tx.Debug().Where("id_faq = ?", id).First(&entity.FAQ{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Where("id_faq = ?", id).Delete(&entity.FAQ{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

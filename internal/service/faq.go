package service

import (
	"fmt"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
)

type IFAQService interface {
	CreateFAQ(FAQReq *model.CreateFAQ) (*entity.FAQ, error)
	GetFAQ() ([]*entity.FAQ, error)
	UpdateFAQ(id int, faqRequest *model.UpdateFAQ) (*entity.FAQ, error)
	DeleteFAQ(id int) error
}

type FAQService struct {
	FAQRepository repository.IFAQRepository
}

func NewFAQService(FAQRepository repository.IFAQRepository) IFAQService {
	return &FAQService{FAQRepository}
}

func (fs *FAQService) CreateFAQ(FAQReq *model.CreateFAQ) (*entity.FAQ, error) {
	faq := &entity.FAQ{
		Judul_FAQ: FAQReq.Judul_FAQ,
		Isi_FAQ:   FAQReq.Isi_FAQ,
	}

	faq, err := fs.FAQRepository.CreateFAQ(faq)
	if err != nil {
		return nil, err
	}
	return faq, nil
}

func (fs *FAQService) GetFAQ() ([]*entity.FAQ, error) {
	faq, err := fs.FAQRepository.GetFAQ()
	if err != nil {
		return nil, err
	}

	return faq, nil
}

func (fs *FAQService) UpdateFAQ(id int, faqRequest *model.UpdateFAQ) (*entity.FAQ, error) {
	faq, err := fs.FAQRepository.UpdateFAQ(id, faqRequest)
	if err != nil {
		return nil, err
	}
	return faq, nil
}

func (fs *FAQService) DeleteFAQ(id int) error {
	if err := fs.FAQRepository.DeleteFAQ(id); err != nil {
		return fmt.Errorf("failed to delete news: %w", err)
	}

	return nil
}

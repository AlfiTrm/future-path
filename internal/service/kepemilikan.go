package service

import (
	"future-path/entity"
	"future-path/internal/repository"
)

type IKepemilikanService interface {
	GetKepemilikan() ([]*entity.Kepemilikan, error)
}

type KepemilikanService struct {
	KepemilikanRepository repository.IKepemilikanRepository
}

func NewKepemilikanService(KepemilikanRepository repository.IKepemilikanRepository) IKepemilikanService {
	return &KepemilikanService{KepemilikanRepository}
}

func (ks *KepemilikanService) GetKepemilikan() ([]*entity.Kepemilikan, error) {
	kepemilikan, err := ks.KepemilikanRepository.GetKepemilikan()
	if err != nil {
		return nil, err
	}
	return kepemilikan, nil
}

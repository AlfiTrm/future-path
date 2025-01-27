package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type IKepemilikanRepository interface {
	GetKepemilikan() ([]*entity.Kepemilikan, error)
}

type KepemilikanRepository struct {
	db *gorm.DB
}

func NewKepemilikanRepository(db *gorm.DB) IKepemilikanRepository {
	return &KepemilikanRepository{db}
}

func (kr *KepemilikanRepository) GetKepemilikan() ([]*entity.Kepemilikan, error) {
	var kepemilikan []*entity.Kepemilikan
	err := kr.db.Debug().Find(&kepemilikan).Error
	if err != nil {
		return nil, err
	}
	return kepemilikan, nil
}

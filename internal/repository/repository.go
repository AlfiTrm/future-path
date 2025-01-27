package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository        IUserRepository
	BeritaRepository      IBeritaRepository
	SekolahRepository     ISekolahRepository
	UniversitasRepository IUniversitasRepository
	FAQRepository         IFAQRepository
	KepemilikanRepository IKepemilikanRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:        NewUserRepository(db),
		BeritaRepository:      NewBeritaRepository(db),
		SekolahRepository:     NewSekolahRepository(db),
		UniversitasRepository: NewUniversitasRepository(db),
		FAQRepository:         NewFAQRepository(db),
		KepemilikanRepository: NewKepemilikanRepository(db),
	}
}

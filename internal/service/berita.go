package service

import (
	"fmt"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
)

type IBeritaService interface {
	CreateBerita(beritaReq *model.CreateBerita) (*entity.Berita, error)
	GetBeritaSingkat(page int) ([]*entity.Berita, error, int64)
	GetBeritaFull(id int) (*entity.Berita, error)
	UpdateBerita(id int, beritaRequest *model.UpdateBerita) (*entity.Berita, error)
	DeleteBerita(id int) error
}

type BeritaService struct {
	BeritaRepository repository.IBeritaRepository
}

func NewBeritaService(BeritaRepository repository.IBeritaRepository) IBeritaService {
	return &BeritaService{BeritaRepository}
}

func (bs *BeritaService) CreateBerita(beritaReq *model.CreateBerita) (*entity.Berita, error) {
	berita := &entity.Berita{
		Judul_Berita: beritaReq.Judul_Berita,
		Isi_Berita:   beritaReq.Isi_Berita,
	}

	berita, err := bs.BeritaRepository.CreateBerita(berita)
	if err != nil {
		return nil, err
	}
	return berita, nil
}

func (bs *BeritaService) GetBeritaSingkat(page int) ([]*entity.Berita, error, int64) {
	limit := 4
	offset := (page - 1) * limit

	berita, err := bs.BeritaRepository.GetBeritaSingkat(limit, offset)
	if err != nil {
		return nil, err, 0
	}

	totalData, err := bs.BeritaRepository.CountAllBerita()
	if err != nil {
		return nil, err, 0
	}

	return berita, nil, totalData
}

func (bs *BeritaService) GetBeritaFull(id int) (*entity.Berita, error) {
	berita, err := bs.BeritaRepository.GetBeritaFull(id)
	if err != nil {
		return nil, err
	}

	return berita, nil
}

func (bs *BeritaService) UpdateBerita(id int, beritaRequest *model.UpdateBerita) (*entity.Berita, error) {
	berita, err := bs.BeritaRepository.UpdateBerita(id, beritaRequest)
	if err != nil {
		return nil, err
	}
	return berita, nil
}

func (bs *BeritaService) DeleteBerita(id int) error {
	if err := bs.BeritaRepository.DeleteBerita(id); err != nil {
		return fmt.Errorf("failed to delete news: %w", err)
	}

	return nil
}

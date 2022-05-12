package service

import (
	"chen/pkg/dto"
	"chen/pkg/model"
	"chen/pkg/repository"
)

type HallService interface {
	FindAll() ([]model.Hall, error)
	FindByID(id int) (model.Hall, error)
	Create(data dto.HallCreate) error
	Update(id int, data dto.HallUpdate) error
	Delete(id int) error
}

type hallService struct {
	hallRepository repository.HallRepository
}

func NewHallService(repo repository.HallRepository) HallService {
	return &hallService{
		hallRepository: repo,
	}
}

func (hs hallService) FindAll() ([]model.Hall, error) {
	return hs.hallRepository.FindAll()
}

func (hs hallService) FindByID(id int) (model.Hall, error) {
	return hs.hallRepository.FindByID(id)
}

func (hs hallService) Create(data dto.HallCreate) error {
	return hs.hallRepository.Create(data)
}

func (hs hallService) Update(id int, data dto.HallUpdate) error {
	return hs.hallRepository.Update(id, data)
}

func (hs hallService) Delete(id int) error {
	return hs.hallRepository.Delete(id)
}

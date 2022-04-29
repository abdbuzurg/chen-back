package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type HallService interface {
	FindAll() ([]model.Hall, error)
	FindById(id int) (model.Hall, error)
	Create(data dto.HallCreateDTO) error
	Update(id int, data dto.HallUpdateDTO) error
	Delete(id int) error
}

type hallService struct {
	hallRepository repository.HallRepository
}

func NewHallService(repo repository.HallRepository) HallService {
	return hallService{
		hallRepository: repo,
	}
}

func (hs hallService) FindAll() ([]model.Hall, error) {
	return hs.hallRepository.FindAll()
}

func (hs hallService) FindById(id int) (model.Hall, error) {
	return hs.hallRepository.FindById(id)
}

func (hs hallService) Create(data dto.HallCreateDTO) error {
	return hs.hallRepository.Create(data)
}

func (hs hallService) Update(id int, data dto.HallUpdateDTO) error {
	return hs.hallRepository.Update(id, data)
}

func (hs hallService) Delete(id int) error {
	return hs.hallRepository.Delete(id)
}

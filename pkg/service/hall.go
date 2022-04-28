package service

import (
	"chen/model"
	"chen/pkg/repository"
)

type HallService interface {
	FindAll() ([]model.Hall, error)
	FindById(id int) (model.Hall, error)
	Create(data model.HallCreateData) error
	Update(id int, data model.HallUpdateData) error
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

func (hs hallService) Create(data model.HallCreateData) error {
	return hs.hallRepository.Create(data)
}

func (hs hallService) Update(id int, data model.HallUpdateData) error {
	return hs.hallRepository.Update(id, data)
}

func (hs hallService) Delete(id int) error {
	return hs.hallRepository.Delete(id)
}

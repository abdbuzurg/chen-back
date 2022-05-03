package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type HallService interface {
	Find(id int) ([]model.Hall, error)
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

func (hs hallService) Find(id int) ([]model.Hall, error) {
	if id == 0 {
		return hs.hallRepository.FindAll()
	}

	hall, err := hs.hallRepository.FindById(id)
	return append([]model.Hall{}, hall), err
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

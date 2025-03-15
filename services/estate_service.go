package services

import (
	"errors"
	"sawittree/models"
	"sawittree/repositories"
)

type EstateService struct {
	Repo *repositories.EstateRepository
}

func (s *EstateService) CreateEstate(width, length int) (*models.Estate, error) {
	if width < 1 || length < 1 || width > 50000 || length > 50000 {
		return nil, errors.New("invalid estate dimensions")
	}

	estate := &models.Estate{
		Width:  width,
		Length: length,
	}

	err := s.Repo.CreateEstate(estate)
	if err != nil {
		return nil, err
	}

	return estate, nil
}

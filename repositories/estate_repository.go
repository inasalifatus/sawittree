package repositories

import (
	"sawittree/models"

	"gorm.io/gorm"
)

type EstateRepository struct {
	DB *gorm.DB
}

func (repo *EstateRepository) CreateEstate(estate *models.Estate) error {
	return repo.DB.Create(estate).Error
}

func (repo *EstateRepository) GetEstateByID(id string) (*models.Estate, error) {
	var estate models.Estate
	err := repo.DB.Preload("Trees").First(&estate, "id = ?", id).Error
	return &estate, err
}

func (repo *EstateRepository) UpdateEstate(estate *models.Estate) error {
	return repo.DB.Save(estate).Error
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Estate struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Width  int       `json:"width"`
	Length int       `json:"length"`
	Trees  []Tree    `gorm:"foreignKey:EstateID"`
}

func (e *Estate) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

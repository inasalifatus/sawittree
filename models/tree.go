package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tree struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	EstateID uuid.UUID `gorm:"type:uuid" json:"estate_id"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
	Height   int       `json:"height"`
}

func (t *Tree) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

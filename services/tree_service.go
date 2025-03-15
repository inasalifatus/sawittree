package services

import (
	"errors"
	"sawittree/models"
	"sawittree/repositories"

	"github.com/google/uuid"
)

type TreeService struct {
	Repo *repositories.EstateRepository
}

// AddTree menambahkan pohon ke dalam estate berdasarkan ID
func (s *TreeService) AddTree(estateID string, x, y, height int) (*models.Tree, error) {
	// Validasi tinggi pohon
	if height < 1 || height > 30 {
		return nil, errors.New("tree height must be between 1 and 30 meters")
	}

	// Ambil estate dari repository
	estate, err := s.Repo.GetEstateByID(estateID)
	if err != nil {
		return nil, errors.New("estate not found")
	}

	// Validasi koordinat pohon
	if x < 1 || x > estate.Width || y < 1 || y > estate.Length {
		return nil, errors.New("tree coordinates out of bounds")
	}

	// Cek apakah sudah ada pohon di koordinat tersebut
	for _, tree := range estate.Trees {
		if tree.X == x && tree.Y == y {
			return nil, errors.New("plot already has a tree")
		}
	}

	// Buat pohon baru
	newTree := models.Tree{
		ID:     uuid.New(),
		X:      x,
		Y:      y,
		Height: height,
	}

	// Tambahkan pohon ke estate dan simpan
	estate.Trees = append(estate.Trees, newTree)
	err = s.Repo.UpdateEstate(estate)
	if err != nil {
		return nil, errors.New("failed to save tree data")
	}

	return &newTree, nil
}

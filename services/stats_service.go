package services

import (
	"errors"
	"sort"
	"sawittree/models"
	"sawittree/repositories"
)

type StatsService struct {
	Repo *repositories.EstateRepository
}

// GetStats mengambil statistik pohon dalam estate
func (s *StatsService) GetStats(estateID string) (map[string]int, error) {
	estate, err := s.Repo.GetEstateByID(estateID)
	if err != nil {
		return nil, errors.New("estate not found")
	}

	trees := estate.Trees
	count := len(trees)

	if count == 0 {
		return map[string]int{
			"count":  0,
			"max":    0,
			"min":    0,
			"median": 0,
		}, nil
	}

	// Menghitung min, max, dan median tinggi pohon
	minHeight, maxHeight, medianHeight := calculateTreeStats(trees)

	return map[string]int{
		"count":  count,
		"max":    maxHeight,
		"min":    minHeight,
		"median": medianHeight,
	}, nil
}

// Fungsi untuk menghitung statistik pohon (min, max, median)
func calculateTreeStats(trees []models.Tree) (int, int, int) {
	heights := []int{}
	for _, tree := range trees {
		heights = append(heights, tree.Height)
	}

	// Sort untuk mendapatkan min, max, dan median
	sort.Ints(heights)

	minHeight := heights[0]
	maxHeight := heights[len(heights)-1]
	medianHeight := calculateMedian(heights)

	return minHeight, maxHeight, medianHeight
}

// Fungsi untuk menghitung median dari daftar tinggi pohon
func calculateMedian(heights []int) int {
	n := len(heights)
	if n%2 == 1 {
		return heights[n/2]
	}
	return (heights[n/2-1] + heights[n/2]) / 2
}

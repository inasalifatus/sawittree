package services

import (
	"errors"
	"fmt"
	"sawittree/models"
	"sawittree/repositories"
)

type DroneService struct {
	Repo *repositories.EstateRepository
}

func (s *DroneService) CalculateDronePath(id string) (int, error) {
	estate, err := s.Repo.GetEstateByID(id)
	if err != nil {
		return 0, errors.New("estate not found")
	}

	width, length := estate.Width, estate.Length
	trees := make(map[string]int)

	for _, tree := range estate.Trees {
		key := keyForPlot(tree.X, tree.Y)
		trees[key] = tree.Height
	}

	totalDistance := 0
	prevHeight := 0

	for y := 1; y <= length; y++ {
		for x := 1; x <= width; x++ {
			if y%2 == 1 { // Move east
				totalDistance += 10
			} else { // Move west
				totalDistance -= 10
			}

			currentHeight := trees[keyForPlot(x, y)]
			verticalMove := abs(currentHeight-prevHeight) + 1
			totalDistance += verticalMove
			prevHeight = currentHeight
		}
	}

	// Descend to ground at the end
	totalDistance += prevHeight

	return totalDistance, nil
}

func keyForPlot(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// CalculateDronePathWithLimit menghitung total jarak dengan batas maksimal jarak drone
func (s *DroneService) CalculateDronePathWithLimit(estateID string, maxDistance int) (int, int, int, error) {
	estate, err := s.Repo.GetEstateByID(estateID)
	if err != nil {
		return 0, 0, 0, errors.New("estate not found")
	}

	totalDistance, restX, restY := calculateDroneDistanceWithLimit(*estate, maxDistance)
	return totalDistance, restX, restY, nil
}

// Function untuk menghitung jarak drone dengan batas maksimum
func calculateDroneDistanceWithLimit(estate models.Estate, maxDistance int) (int, int, int) {
	width := estate.Width
	length := estate.Length
	trees := estate.Trees

	horizontalDistance := 0
	verticalDistance := 0
	verticalLiftDistance := 0
	totalDistance := 0

	restX, restY := 1, 1

	for y := 1; y <= length; y++ {
		if y%2 == 1 {
			horizontalDistance += (width - 1)
		} else {
			horizontalDistance += (width - 1)
		}

		verticalDistance++
		verticalLiftDistance += calculateVerticalLiftDistanceAtRow(trees, y)

		totalDistance = (horizontalDistance + verticalDistance + verticalLiftDistance) * 10

		// Jika total jarak melebihi batas, return koordinat terakhir sebelum kehabisan baterai
		if totalDistance >= maxDistance {
			return maxDistance, restX, restY
		}

		restX = width
		restY = y
	}

	return totalDistance, restX, restY
}

// Function untuk menghitung jarak naik-turun pada setiap baris
func calculateVerticalLiftDistanceAtRow(trees []models.Tree, row int) int {
	totalLift := 0
	for _, tree := range trees {
		if tree.Y == row {
			totalLift += tree.Height + 1
		}
	}
	return totalLift
}

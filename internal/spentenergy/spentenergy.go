package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count = 0 or less")
	}
	if weight <= 0 {
		return 0, errors.New("weight = 0 or less")
	}
	if height <= 0 {
		return 0, errors.New("height = 0 or less")
	}
	if duration <= 0 {
		return 0, errors.New("walk duration = 0 or less")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count = 0 or less")
	}
	if weight <= 0 {
		return 0, errors.New("weight = 0 or less")
	}
	if height <= 0 {
		return 0, errors.New("height = 0 or less")
	}
	if duration <= 0 {
		return 0, errors.New("walk duration = 0 or less")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * meanSpeed * durationInMinutes) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}

	if height <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	totalWalkingDistance := (float64(steps) * stepLength) / mInKm

	return totalWalkingDistance
}

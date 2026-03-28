package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	ErrInvalidStepsCount = errors.New("invalid steps count")
	ErrInvalidDuration   = errors.New("invalid duration")
	ErrInvalidHeight     = errors.New("invalid height")
	ErrInvalidWeight     = errors.New("invalid weight")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", ErrInvalidStepsCount)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", ErrInvalidHeight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", ErrInvalidWeight)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", ErrInvalidDuration)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	return (weight * meanSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	length := stepLength * float64(steps)
	return length / mInKm
}

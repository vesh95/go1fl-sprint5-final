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
	errInvalidStepsCount = errors.New("invalid steps count")
	errInvalidDuration   = errors.New("invalid duration")
	errInvalidHeight     = errors.New("invalid height")
	errInvalidWeight     = errors.New("invalid weight")
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
		return 0, fmt.Errorf("running spent calories: %w", errInvalidStepsCount)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", errInvalidWeight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", errInvalidHeight)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("running spent calories: %w", errInvalidDuration)
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

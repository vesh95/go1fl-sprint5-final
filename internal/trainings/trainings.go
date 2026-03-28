package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	TrainingTypeRunning string = "Бег"
	TrainingTypeWalking string = "Ходьба"
)

var (
	ErrInvalidParam      = errors.New("invalid params string")
	ErrInvalidStepsCount = errors.New("invalid steps count")
	ErrInvalidDuration   = errors.New("invalid duration")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	params := strings.Split(datastring, ",")
	if len(params) != 3 {
		return fmt.Errorf("parse training error: %w", ErrInvalidParam)
	}
	t.Steps, err = strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("parse training error: %w", ErrInvalidStepsCount)
	}
	t.TrainingType = params[1]
	t.Duration, err = time.ParseDuration(params[2])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("parse training error: %w", ErrInvalidDuration)
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	calories := 0.0
	var err error
	switch t.TrainingType {
	case TrainingTypeRunning:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case TrainingTypeWalking:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}

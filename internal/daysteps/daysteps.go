package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrInvalidParam      = errors.New("invalid params string")
	ErrInvalidStepsCount = errors.New("invalid steps count")
	ErrInvalidDuration   = errors.New("invalid duration")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	params := strings.Split(datastring, ",")
	if len(params) != 2 {
		return fmt.Errorf("parse training error: %w", ErrInvalidParam)
	}
	ds.Steps, err = strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("parse training error: %w", ErrInvalidStepsCount)
	}
	ds.Duration, err = time.ParseDuration(params[1])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("parse training error: %w", ErrInvalidDuration)
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("daysteps info error: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}

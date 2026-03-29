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
	errInvalidParam      = errors.New("invalid params string")
	errInvalidStepsCount = errors.New("invalid steps count")
	errInvalidDuration   = errors.New("invalid duration")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	params := strings.Split(datastring, ",")
	if len(params) != 2 {
		return fmt.Errorf("parse training error: %w", errInvalidParam)
	}
	ds.Steps, err = strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("parse training error: %w", errInvalidStepsCount)
	}
	ds.Duration, err = time.ParseDuration(params[1])
	if err != nil {
		return fmt.Errorf("parse training error: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("parse training error: %w", errInvalidDuration)
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

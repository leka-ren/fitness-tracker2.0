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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitedDatastring := strings.Split(datastring, ",")

	if len(splitedDatastring) != 2 {
		return errors.New("incorrect arguments count, should be 2")
	}

	steps, err := strconv.Atoi(splitedDatastring[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("steps count is 0 or less")
	}

	ds.Steps = steps

	durationTime, err := time.ParseDuration(splitedDatastring[1])
	if err != nil {
		return err
	}

	if durationTime.Seconds() <= 0 {
		return errors.New("time duration is 0 or less")
	}

	ds.Duration = durationTime

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)
	return info, nil
}

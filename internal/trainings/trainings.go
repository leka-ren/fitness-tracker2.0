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

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitedDatastring := strings.Split(datastring, ",")
	if len(splitedDatastring) != 3 {
		return errors.New("incorrect arguments count, should be 3")
	}

	steps, err := strconv.Atoi(splitedDatastring[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("steps count is 0 or less")
	}

	t.Steps = steps

	if splitedDatastring[1] == "" {
		return errors.New("trainig type is empty")
	}

	t.TrainingType = splitedDatastring[1]

	durationTime, err := time.ParseDuration(splitedDatastring[2])
	if err != nil {
		return err
	}

	if durationTime.Seconds() <= 0 {
		return errors.New("time duration is 0 or less")
	}

	t.Duration = durationTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var spentCaloriesInfo string
	switch t.TrainingType {
	case "Бег":
		res, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

		spentCaloriesInfo = fmt.Sprintf("Сожгли калорий: %.2f", res)
	case "Ходьба":
		res, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}
		spentCaloriesInfo = fmt.Sprintf("Сожгли калорий: %.2f", res)
	default:
		return "", errors.New("unknown training type")
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\n%s\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), spentCaloriesInfo)

	return info, nil
}

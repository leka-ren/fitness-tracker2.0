package trainings

import (
	"errors"
	"fmt"
	"log"
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
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitedDatastring := strings.Split(datastring, ",")
	if len(splitedDatastring) != 3 {
		return errors.New("wrong paramets length")
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
		log.Println(splitedDatastring[2], durationTime, durationTime.Seconds())
		return errors.New("time duration is 0 or less")
	}

	t.Duration = durationTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	info := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	info += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	info += fmt.Sprintf("Дистанция: %.2f км.\n", spentenergy.Distance(t.Steps, t.Personal.Height))
	info += fmt.Sprintf("Скорость: %.2f км/ч\n", spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration))
	switch t.TrainingType {
	case "Бег":
		res, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)

		if err != nil {
			return "", err
		}

		info += fmt.Sprintf("Сожгли калорий: %.2f\n", res)
	case "Ходьба":
		res, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)

		if err != nil {
			return "", err
		}

		info += fmt.Sprintf("Сожгли калорий: %.2f\n", res)
	default:
		return "", errors.New("unknown training type")
	}
	return info, nil
}

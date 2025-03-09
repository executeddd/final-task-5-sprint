package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("ошибка, слайс должен равняться трём")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании шагов: %w", err)
	}
	t.Steps = steps
	if parts[1] != "Бег" && parts[1] != "Ходьба" {
		return errors.New("ошибка, неизвестный тип активности")
	}
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании длительности активности %w", err)
	}
	t.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", errors.New("ошибка, длительность равна или меньше нуля")
	}
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var spentCalories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("\n Тип тренировки: %s\n Длительность: %.2f\n Дистанция: %.2f\n Скорость: %.2f\n Сожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories), nil
}

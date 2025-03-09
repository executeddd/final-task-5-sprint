package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("длина слайса должна равняться двум")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании количества шагов")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании длительности активности")
	}
	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds *DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("ошибка, длительность активности равна или меньше нуля")
	}
	distance := spentenergy.Distance(ds.Steps)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка при подсчёте сожженых калорий")
	}
	return fmt.Sprintf("\nКоличество шагов: %d.\n Дистанция составила %.2f км.\n Вы сожгли %.2f ккал.",
		ds.Steps, distance, spentCalories), nil
}

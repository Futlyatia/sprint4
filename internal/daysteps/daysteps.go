package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	objects := strings.Split(data, ",")
	if len(objects) != 2 {
		return 0, 0, errors.New("неверный формат - неправильное количество параметров")
	}
	if len(objects) == 0 {
		return 0, 0, errors.New("пустой ввод")
	}
	steps, err := strconv.Atoi(objects[0])
	if err != nil {
		return 0, 0, errors.New("некорректное количество шагов")
	}
	duration, err := time.ParseDuration(objects[1])
	if err != nil {
		return 0, 0, errors.New("некорректная продолжительность")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float32(steps) * stepLength
	distance /= mInKm
	cals, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		return ""
	}
	//fmt.Println(cals)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, cals)
}

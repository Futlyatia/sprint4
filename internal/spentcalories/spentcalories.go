package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	objects := strings.Split(data, ",")
	if len(objects) != 3 {
		return 0, "", 0, errors.New("некорректный формат данных")
	}
	steps, err := strconv.Atoi(objects[0])
	if err != nil {
		return 0, "", 0, errors.New("некорректное количество шагов")
	}
	duration, err := time.ParseDuration(objects[2])
	if err != nil {
		return 0, "", 0, errors.New("некорректная продолжительность")
	}
	return steps, objects[1], duration, nil
}

func distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	dist := float64(steps) * lenStep
	dist /= mInKm
	return float64(dist)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	speed := dist / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	switch training {
	case "Бег":
		cals, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		response := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			training, duration.Hours(), dist, speed, cals)
		return response, nil
	case "Ходьба":
		cals, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		response := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			training, duration.Hours(), dist, speed, cals)
		return response, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, errors.New("отрицательные шаги")
	}
	if steps == 0 {
		return 0, errors.New("нулевые шаги")
	}
	if weight < 0 {
		return 0, errors.New("отрицательный вес")
	}
	if weight == 0 {
		return 0, errors.New("нулевой вес")
	}
	if height == 0 {
		return 0, errors.New("нулевой рост")
	}
	if height < 0 {
		return 0, errors.New("отрицательный рост")
	}
	if duration < 0 {
		return 0, errors.New("отрицательная продолжительность")
	}
	if duration == 0 {
		return 0, errors.New("нулевая продолжительность")
	}
	speed := meanSpeed(steps, height, duration)
	cals := (weight * speed * duration.Minutes()) / minInH
	return cals, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, errors.New("отрицательные шаги")
	}
	if steps == 0 {
		return 0, errors.New("нулевые шаги")
	}
	if weight < 0 {
		return 0, errors.New("отрицательный вес")
	}
	if weight == 0 {
		return 0, errors.New("нулевой вес")
	}
	if height == 0 {
		return 0, errors.New("нулевой рост")
	}
	if height < 0 {
		return 0, errors.New("отрицательный рост")
	}
	if duration < 0 {
		return 0, errors.New("отрицательная продолжительность")
	}
	if duration == 0 {
		return 0, errors.New("нулевая продолжительность")
	}
	speed := meanSpeed(steps, height, duration)
	cals := (weight * speed * duration.Minutes()) / minInH
	return cals * walkingCaloriesCoefficient, nil
}

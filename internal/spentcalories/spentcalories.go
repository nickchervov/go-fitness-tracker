package spentcalories

import (
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
	// TODO: реализовать функцию

	datas := strings.Split(data, ",")
	if len(datas) != 3 {
		return 0, "", 0, fmt.Errorf("во входящих данных недостаточно аргументов.")
	}

	steps, err := strconv.Atoi(datas[0])
	if err != nil {
		return 0, "", 0, err
	}

	action := datas[1]

	duration, err := time.ParseDuration(datas[2])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 || duration <= 0 {
		return 0, "", 0, fmt.Errorf("введены некорректные значения.")
	}

	return steps, action, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepLength := height * stepLengthCoefficient

	distanceInM := stepLength * float64(steps)
	distanceInKM := distanceInM / mInKm

	return distanceInKM
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 {
		return 0
	}

	fullDistance := distance(steps, height)

	averageSpeed := fullDistance / duration.Hours()

	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	output := `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`

	steps, action, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if steps <= 0 || duration <= 0 {
		log.Println("введены некорректные значения.")
	}

	switch action {
	case "Бег":
		fullDistance := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, duration)
		spentCalories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(output, action, duration.Hours(), fullDistance, averageSpeed, spentCalories), nil
	case "Ходьба":
		fullDistance := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, duration)
		spentCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(output, action, duration.Hours(), fullDistance, averageSpeed, spentCalories), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки.")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("введены некорректные значения.")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMin := duration.Minutes()

	spentCalories := (weight * averageSpeed * durationInMin) / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("введены некорректные значения.")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMin := duration.Minutes()

	spentCalories := (weight * averageSpeed * durationInMin) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}

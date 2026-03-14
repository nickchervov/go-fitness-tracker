package daysteps

import (
	"fmt"
	"log"
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
	// TODO: реализовать функцию
	datas := strings.Split(data, ",")
	if len(datas) != 2 {
		return 0, 0, fmt.Errorf("not enough arguments in the incoming data.")
	}

	steps, err := strconv.Atoi(datas[0])
	if err != nil {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(datas[1])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps less than or equal to 0.")
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration less than or equal to 0.")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	output := `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		log.Println("steps less than or equal to 0.")
		return ""
	}

	if duration <= 0 {
		log.Println("duration less than or equal to 0.")
		return ""
	}

	distanceInM := float64(steps) * stepLength
	distanceInKM := distanceInM / mInKm

	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf(output, steps, distanceInKM, spentCalories)
}

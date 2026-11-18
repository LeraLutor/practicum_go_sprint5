package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Расчет дистанции в км
func Distance(steps int, height float64) float64 {
	// подсчет длины шага
	stepLength := height * stepLengthCoefficient

	// подсчет дистанции и возврат значения
	return float64(steps) * stepLength / mInKm
}

// Расчет средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка шагов, роста и длительности тренировки на корректность
	if steps < 0 || height <= 0 || duration <= 0 {
		return 0
	}

	// подсчет средней скорости и возврат значения
	return Distance(steps, height) / duration.Hours()
}

// Расчет ккал потраченных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка значения колва шагов
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}

	// проверка значения веса
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}

	// проверка значения роста
	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}

	// проверка значения длительности тренировки
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	// расчет средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// подсчет колва калорий
	calories := weight * meanSpeed * duration.Minutes() / minInH

	return calories, nil
}

// Расчет ккал потраченных при хотьбн
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка значения колва шагов
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}

	// проверка значения веса
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}

	// проверка значения роста
	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}

	// проверка значения длительности тренировки
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	// расчет средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// подсчет колва калорий
	calories := (weight * meanSpeed * duration.Minutes() / minInH) * walkingCaloriesCoefficient

	return calories, nil
}

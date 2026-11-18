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

// Структура данных о тренировке
type Training struct {
	Steps                 int           // колво шагов
	TrainingType          string        // тип тренировки
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // данные о человеке
}

// Парсинг данных
func (t *Training) Parse(datastring string) (err error) {
	// разделение данных на слайс
	parts := strings.Split(datastring, ",")
	// проверка на длину входных данных
	if len(parts) != 3 {
		return errors.New(
			"некорректный формат тренировки: ожидается шаги,тип,продолжительность",
		)
	}

	// преобразование шагов в int
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	// проверка на ошибку преобразования
	if err != nil {
		return fmt.Errorf("не удалось распознать количество шагов: %w", err)
	}
	// проверка на отрицательное колво шагов
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	// запись в структуру
	t.Steps = steps

	// преобразование типа тренировки
	trainingType := strings.TrimSpace(parts[1])
	// запись в структуру
	t.TrainingType = trainingType

	// преобразование времени тренировки
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	// проверка на ошибку преобразования
	if err != nil {
		return fmt.Errorf(
			"не удалось распознать продолжительность тренировки: %w",
			err,
		)
	}
	// проверка на отрицательное время
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	// запись в структуру
	t.Duration = duration

	return nil
}

// Формирование данных о тренировке
func (t Training) ActionInfo() (string, error) {
	// расчет дистанции
	distance := spentenergy.Distance(t.Steps, t.Height)
	// расчет скорости
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// переменные для структуру тренировки
	var calories float64
	var err error

	// расчет ккал исходя из вида тренировки
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	// ошибка для неизвестного типа тренировки
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// проверка на ошибку
	if err != nil {
		return "", err
	}

	// вывод инфо о тренировке
	info := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return info, nil
}

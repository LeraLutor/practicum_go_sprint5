package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Структура данных о шагах
type DaySteps struct {
	Steps                 int           // колво шагов
	Duration              time.Duration // время тренировки
	personaldata.Personal               // данные о человеке
}

// Парсинг данных
func (ds *DaySteps) Parse(datastring string) (err error) {
	// разделение данных на слайс
	parts := strings.Split(datastring, ",")
	// проверка на длину входных данных
	if len(parts) != 2 {
		return errors.New(
			"некорректный формат прогулки: ожидается шаги,продолжительность",
		)
	}

	// преобразование шагов
	steps, err := strconv.Atoi(parts[0])
	// проверка на ошибку преобразования
	if err != nil {
		return fmt.Errorf("не удалось распознать количество шагов: %w", err)
	}
	// проверка на отрицательное колво шагов
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	// запись в структуру
	ds.Steps = steps

	// преобразование времени тренировки
	duration, err := time.ParseDuration(parts[1])
	// проверка на ошибку преобразования
	if err != nil {
		return fmt.Errorf(
			"не удалось распознать продолжительность прогулки: %w",
			err,
		)
	}
	// проверка на отрицательное время
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	// запись в структуру
	ds.Duration = duration

	return nil
}

// Формирование данных о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	// расчет дистанции
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// расчет сожженых ккал
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	// проверка на ошибку
	if err != nil {
		return "", err
	}

	// вывод инфо о тренировке
	info := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return info, nil
}

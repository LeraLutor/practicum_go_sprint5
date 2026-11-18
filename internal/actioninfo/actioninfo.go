package actioninfo

import (
	"fmt"
	"log"
)

// Интерфейс с сигнатурами
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Вывод информации об активности
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// проверка на ошибку парсинга
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		// формирования инфо активности
		info, err := dp.ActionInfo()
		// проаерка на ошибку и логирование
		if err != nil {
			log.Printf(
				"ошибка формирования информации для %q: %v",
				data,
				err,
			)
			continue
		}

		// вывод информации
		fmt.Println(info)
	}
}

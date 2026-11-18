package personaldata

import (
	"fmt"
)

// Структура данных пользователя
type Personal struct {
	Name   string  // имя
	Weight float64 // вес
	Height float64 // рост
}

// Вывод данных пользователя
func (p Personal) Print() {
	fmt.Printf(
		"Имя: %s\nВес: %.2f\nРост: %.2f\n",
		p.Name,
		p.Weight,
		p.Height,
	)
}

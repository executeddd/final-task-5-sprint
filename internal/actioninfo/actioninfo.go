package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Println("\nошибка парсинга данных: ", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("\nошибка при формировании строки информации")
		}
		fmt.Println(info)
	}
}

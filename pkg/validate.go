package pkg

import "regexp"

func CityValidate(city string) bool {
	// Регулярное выражение, проверяющее, содержит ли строка только буквы
	regex := regexp.MustCompile("^[a-zA-Z]+$")

	// Сопоставляем строку с регулярным выражением
	return regex.MatchString(city)
}

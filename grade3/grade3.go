package grade3

import (
	"fmt"
	"regexp"
	"strings"
)

func Task41() {
	products := []map[string]interface{}{
		{"id": 1001, "name": "Product 1"},
		{"id": 1002, "name": "Product 2"},
		{"id": 1003, "name": "Product 3"},
		{"id": 1006, "name": "Product 4"},
	}

	productMap := make(map[int]string)
	for _, product := range products {
		id := product["id"].(int)
		name := product["name"].(string)
		productMap[id] = name
	}

	id := 1006
	if name, exists := productMap[id]; exists {
		println("Product name:", name)
	} else {
		println("Product not found")
	}
}

func Task42(s string) string {
	parts := stringToParts(s)
	for i := 1; i < len(parts); i += 2 {
		parts[i] = reverseString(parts[i])
	}
	return strings.Join(parts, "")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stringToParts(s string) []string {
	partLength := 2 // Это значение можно варьировать для увеличения сложности шифрования
	parts := make([]string, 0)
	for i := 0; i < len(s); i += partLength {
		end := i + partLength
		if end > len(s) {
			end = len(s)
		}
		parts = append(parts, s[i:end])
	}
	return parts
}

func fields() []string {
	return []string{"id", "name", "email", "age", "created_at"}
}

func requiredFields(requestedFields string) []string {
	modelFields := fields()
	requested := strings.Split(requestedFields, ",")

	fieldMap := make(map[string]bool)
	for _, field := range modelFields {
		fieldMap[field] = true
	}

	var validFields []string
	for _, field := range requested {
		trimmedField := strings.TrimSpace(field)
		if fieldMap[trimmedField] {
			validFields = append(validFields, trimmedField)
		}
	}
	return validFields
}

func Task43() {
	queryFields := "id, name, phone"
	result := requiredFields(queryFields)
	fmt.Println("Required fields:", result)
}

func Task5() {
	name := "Салтыков-Щедрин"
	if validateName(name) {
		fmt.Printf("Имя \"%s\" валидно\n", name)
	} else {
		fmt.Printf("Имя \"%s\" невалидно\n", name)
	}

	code := "code_123-ABC"
	if validateCode(code) {
		fmt.Printf("Символьный код \"%s\" валиден\n", code)
	} else {
		fmt.Printf("Символьный код \"%s\" невалиден\n", code)
	}
}

func validate(input string, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Ошибка компиляции регулярного выражения:", err)
		return false
	}
	return re.MatchString(input)
}

func validateName(name string) bool {
	pattern := `^[a-zA-Zа-яА-ЯёЁ'\-\s]+$`
	return validate(name, pattern)
}

func validateCode(code string) bool {
	pattern := `^[a-zA-Z0-9\-_]+$`
	return validate(code, pattern)
}

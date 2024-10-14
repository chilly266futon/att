package grade3

import (
	"errors"
	"fmt"
	"os"
)

func OpenFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла %s: %w", filename, err)
	}
	return data, nil
}

type Config struct {
	Name  string
	Value int
}

func ReadConfig(filename string) (Config, error) {
	var config Config
	data, err := OpenFile(filename)
	if err != nil {
		return config, fmt.Errorf("ошибка чтения конфигурации: %w", err)
	}
	if _, err := fmt.Sscanf(string(data), "%s %d", &config.Name, &config.Value); err != nil {
		return config, fmt.Errorf("ошибка парсинга конфигурации: %w", err)
	}
	return config, nil
}

type NotFoundError struct {
	Filename string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("файл %s не найден", e.Filename)
}

func IsNotFoundError(err error) bool {
	var notFoundErr *NotFoundError
	return errors.As(err, &notFoundErr)
}

func HandleFile(filename string) {
	_, err := OpenFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Файл не существует:", filename)
		} else {
			fmt.Printf("Ошибка открытия файла: %v\n", err)
		}
		return
	}

	config, err := ReadConfig(filename)
	if err != nil {
		if IsNotFoundError(err) {
			fmt.Println("Конфигурационный файл не найден:", filename)
		} else {
			fmt.Printf("Ошибка чтения конфигурации: %v\n", err)
		}
		return
	}

	fmt.Printf("Успешно прочитана конфигурация: %+v\n", config)
}

func Task6() {
	// существующий файл
	HandleFile("config.txt")

	// несуществующий файл
	HandleFile("missing.txt")

	// неправильный конфиг
	HandleFile("bad_config.txt")
}

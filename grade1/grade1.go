package grade1

import (
	"fmt"
)

func Greeting() error {
	var name string
	if _, err := fmt.Scanln(&name); err != nil {
		return err
	}
	msg := fmt.Sprintf("Привет, %s!", name)
	fmt.Println(msg)
	return nil
}

func Cube() error {
	var n, volume, area int
	if _, err := fmt.Scanln(&n); err != nil {
		return err
	}
	volume = n * n * n
	area = n * n * 6
	fmt.Println("Объем:", volume)
	fmt.Println("Площадь:", area)
	return nil
}

func PrevAndNext() error {
	var num int
	if _, err := fmt.Scanln(&num); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Следующее за числом %d число: %d", num, num+1))
	fmt.Println(fmt.Sprintf("Для числа %d предыдущее число: %d", num, num-1))
	return nil
}

func DigitsOfNum(n int) (int, int, int, int) {
	var a, b, c, d int
	a = n / 1000
	b = n % 1000 / 100
	d = n % 10
	c = n / 10 % 10
	return a, b, c, d
}

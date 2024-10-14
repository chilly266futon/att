package grade2

import (
	"fmt"
	"math"
)

func Circle(r float64) {
	s := math.Pi * r * r
	c := 2 * math.Pi * r

	fmt.Println("S =", s)
	fmt.Println("C =", c)
}

func IsLeap(year int) {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("YES")
				return
			}
			fmt.Println("NO")
			return
		}
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
	return
}

// FactorialR returns the factorial of n using recursion
func FactorialR(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialR(n-1)
}

// FactorialC returns the factorial of n using a cycle
func FactorialC(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

func TriangleOfDigits(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

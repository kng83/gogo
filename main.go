package main

import (
	"fmt"
	"mymodule"
)

func main() {
	mymodule.SayHello()
	liczba := 4.4
	var length float64 = 1.2
	fmt.Println(liczba, int(length))
	repeatLine("this is a test string", 3)

	area, err := countArea(-4.4, 2.2)
	if !err {
		println("area is ", area)
	} else {
		println("we have an error")
	}

	// fmt.Println(reflect.TypeOf(liczba))

	bobo, erra := retError()
	if erra != nil {
		println("thieres is no error")
	} else {
		println("there is s", bobo)
	}

}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

// Count area with error handling
func countArea(width float32, height float32) (float32, bool) {
	if width < 0 || height < 0 {
		return 0.0, true
	}
	return width * height, false
}

func retError() (int, error) {
	return 0, fmt.Errorf("Nie można obliczyć pierwiastka liczby ujemnej.")
}

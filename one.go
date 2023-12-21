package main

import (
	"fmt"
	"math"
	"reflect"
	"shapes"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(shapes.CalcRectArea(2, 3))
	fmt.Println(shapes.GetMax(10, 11))
	fmt.Println(shapes.GetSum())

	today := time.Now().Weekday()
	fmt.Println(today + 1)

	some := 4
	p := &some
	println(p)

	type Vertex struct {
		X int
		Y int
	}

	structField := Vertex{1, 2}
	structField.X = 44

	fmt.Println(Vertex{1, 2})
	fmt.Println(structField.X)

	// array check
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, primes[0], len(primes))

	//slice example
	upperBound := 3
	second := primes[1:upperBound]
	fmt.Println(second)

	//other array
	third := [4]int{1, 2, 3, 4}
	www := third[1:3]
	four := append(www, 22, 32, 11)

	printSlice(four, "slice of four")
	printSlice(www, "slice of www")

	//Check map
	type Point struct {
		Lat  float64
		Long float64
	}

	var m = map[string]Point{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)

	keys := reflect.ValueOf(m).MapKeys()
	fmt.Println(keys)

	// czek to function
	comp1 := comp(10, addOne)
	fmt.Println(comp1, "here is comp1")

	//Sprawdzenie domkniecia
	someAdder := adder()
	fmt.Println(someAdder(), someAdder())

	//Sprawdzenie metod dodanych do struktury
	wierz := Wierzcholek{10, 20}
	fmt.Println(wierz.Abs())

	var dzisJest SuperDay
	fmt.Println(dzisJest.Day(2))

}

// Used to print table
func printSlice(s []int, coment string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
		if i == len(s)-1 {
			fmt.Printf(" %s\n", coment)
		}
	}
}

// Przyklad dwoch funkcji ktore razem wspolpracuja
func comp(value int, fn func(int) int) int {
	return fn(value)
}

func addOne(value int) int {
	return value + 1
}

func adder() func() int {
	sum := 0
	return func() int {
		sum++
		return sum
	}
}

// Wierzcholek Funkcja z metodami do struktury
type Wierzcholek struct {
	X float64
	Y float64
}

// Abs tutaj jest dodawanie metod do vertex
func (v Wierzcholek) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v Wierzcholek) Cos() int {
	return 10
}

type SuperDay int

func (v SuperDay) Day(pick int) string {
	if pick == 1 {
		return "Poniedzialek"
	}

	if pick == 2 {
		return "wtorek"
	}
	return "brak dnia"
}

func AddTWo(value int) int {
	return value + 2
}

type Koko interface {
	AddTwo() int
}

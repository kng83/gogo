package main

import "fmt"

//MyInterface iplementation
type MyInterface interface {
	MyMethod()
	ChangeX()
	PrintX()
	CheckOtherX()
}

//SecondInterface is here
type SecondInterface interface {
	SecondMethod()
	AssignValue(value MyInt)
}

//Tobi implementation
type Tobi struct {
	S string
	X int
}

//MyInt is some type
type MyInt int

//MyMethod description
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t Tobi) MyMethod() {
	fmt.Println(t.S)
}

//ChangeX desc
func (t *Tobi) ChangeX() {
	t.X = 44
}

//CheckOtherX desc
func (t Tobi) CheckOtherX() {
	t.X = 55
}

//PrintX print variable
func (t Tobi) PrintX() {
	fmt.Println(t.X)
}

//SecondMethod is here
func (someInt *MyInt) SecondMethod() {
	if someInt == nil {
		fmt.Println("nil is comming")
	} else {
		fmt.Println("there is normal value")
	}
}

//AssignValue is here
func (someInt *MyInt) AssignValue(value MyInt) {
	*someInt = value
}

func main() {
	// Jesli chcemy zmieniac wartosci w strukturze to musimy dostawac sie przez adres
	var i MyInterface = &Tobi{"some text", 10}
	i.MyMethod()
	i.ChangeX()
	i.CheckOtherX()
	i.PrintX()

	var myInt *MyInt           //deklaracja wskaznika
	var second SecondInterface // tu dajemy interfejs
	second = myInt             // tu przypisujemy myInt do second

	//	second.AssignValue(10)
	second.SecondMethod() //ta metoda bedzie miala nila
}

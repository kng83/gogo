package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// tutaj jest czytanie stringu i na podstawie jego jest budowany buffor ascii
func firstOne() {
	var someText string = "aaaa"
	r := strings.NewReader(someText)

	b := make([]byte, len(someText))
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
func second() {
	//jezeli zwracany jest nil to plik jest dobrze odczytany
	b, err := ioutil.ReadFile("./test_file.txt")

	if err == nil {
		fmt.Printf("%v\n", strconv.FormatInt(int64(b[0]), 16))
		for i := 0; i < len(b); i++ {
			fmt.Printf("%v", b[i])
		}

	} else {
		fmt.Println("this is a err")
	}
}
func main() {
	//firstOne()
	second()
}

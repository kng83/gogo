package shapes

// CalcRectArea function
func CalcRectArea(a int, b int) int {
	return a * b
}

// GetMax from to values
func GetMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetSum of 10 elements
func GetSum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

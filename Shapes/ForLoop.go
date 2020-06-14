package shapes

// GetSum of 10 elements
func GetSum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

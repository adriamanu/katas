package main

// grow is a function that multiply in order each values of the array passed as parameter
func grow(arr []int) int {
	result := 1
	for _, num := range arr {
		result = result * num
	}
	println(result)
	return result
}

func main() {
	grow([]int{1, 2, 3, 4})       // 24
	grow([]int{4, 1, 1, 1, 4})    // 16
	grow([]int{2, 2, 2, 2, 2, 2}) // 64
}

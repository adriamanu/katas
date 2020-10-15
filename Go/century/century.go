package main

func century(year int) int {
	result := year / 100
	if year%100 != 0 {
		result++
	}
	return result
}

func main() {
	century(1705)
	century(1900)
	century(1601)
	century(2000)
}

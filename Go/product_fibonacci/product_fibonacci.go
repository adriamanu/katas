package main

import (
	"fmt"
	"log"
	"time"
)

func fibonacci(n int) uint64 {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func ProductFib(prod uint64) [3]uint64 {
	var i int = 0
	for {
		fi := fibonacci(i)
		fj := fibonacci(i + 1)
		switch {
		case fi*fj == prod:
			return [3]uint64{fi, fj, 1}
		case fi*fj > prod:
			return [3]uint64{fi, fj, 0}
		default:
			i++
		}
	}
}

func main() {

	start := time.Now()
	fmt.Println(ProductFib(715))
	elapsed := time.Since(start)
	log.Printf("ProductFib(715) %s", elapsed)

	start = time.Now()
	fmt.Println(ProductFib(4895))
	elapsed = time.Since(start)
	log.Printf("ProductFib(4895) %s", elapsed)

	start = time.Now()
	fmt.Println(ProductFib(5895))
	elapsed = time.Since(start)
	log.Printf("ProductFib(5895) %s", elapsed)

	start = time.Now()
	fmt.Println(ProductFib(74049690))
	elapsed = time.Since(start)
	log.Printf("ProductFib(74049690) %s", elapsed)

	start = time.Now()
	fmt.Println(ProductFib(84049690))
	elapsed = time.Since(start)
	log.Printf("ProductFib(84049690) %s", elapsed)

}

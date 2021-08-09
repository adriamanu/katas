package main

import (
	"fmt"
	"log"
	"time"
)

func fibonacci(a, b, prod uint64) [3]uint64 {
	f := a + b
	p := b * f
	switch {
	case p == prod:
		return [3]uint64{b, f, 1}
	case p < prod:
		return fibonacci(b, f, prod)
	default:
	}
	return [3]uint64{b, f, 0}
}

func ProductFib(prod uint64) [3]uint64 {
	return fibonacci(1, 2, prod)
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

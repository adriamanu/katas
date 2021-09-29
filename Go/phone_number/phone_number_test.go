package main

import (
	"fmt"
	"log"
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		dpn := "(123) 456-7890"
		pn := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		if pn != dpn {
			log.Fatal(fmt.Sprintf("wrong phone number got %s want %s", pn, dpn))
		}
	})
}

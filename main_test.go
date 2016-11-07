package main

import (
	"fmt"
	"testing"
)

func TestColors(t *testing.T) {

	for i := 0; i < 100; i++ {
		asFloat := float64(i)
		c := tempToRGBA(asFloat)
		fmt.Println(i)
		debugColor(c)
	}

}

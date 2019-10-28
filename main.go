package main

import (
	"fmt"

	"github.com/kortheof/go-slices/compare"
)

func main() {
	s1 := []int{8, 7, 23, 9, 2, 0}
	s2 := []int{0, 23, 8, 9, 2}

	missing := compare.FindMissingElement(s1, s2)
	fmt.Println(missing)
}

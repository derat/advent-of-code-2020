package main

import (
	"fmt"

	"github.com/derat/advent-of-code/lib"
)

func main() {
	const total = 2020
	seen := make(map[int]struct{})
	for _, n := range lib.InputInts("2020/1") {
		rem := total - n
		if _, ok := seen[rem]; ok {
			fmt.Printf("%d + %d = %d, %d * %d = %d\n", n, rem, total, n, rem, n*rem)
		}
		seen[n] = struct{}{}
	}

	for n := range seen {
		r1 := total - n
		for m := range seen {
			r2 := r1 - m
			if _, ok := seen[r2]; ok {
				fmt.Printf("%d + %d + %d = %d, %d * %d * %d = %d\n", n, m, r2, total, n, m, r2, n*m*r2)
			}
		}
	}
}

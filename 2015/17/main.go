package main

import (
	"fmt"

	"github.com/derat/advent-of-code/lib"
)

func main() {
	conts := lib.InputInts("2015/17")
	lib.AssertLessEq(len(conts), 64) // packing into uint64

	const total = 150                    // given in puzzle
	type set map[uint64]struct{}         // container bitfield
	literCombos := make([]set, total+1)  // map from liters to remaining containers
	all := uint64((1 << len(conts)) - 1) // all containers available
	literCombos[0] = set{all: struct{}{}}

	for liters := 1; liters < len(literCombos); liters++ {
		combos := make(set)
		for i, cont := range conts {
			if prev := liters - cont; prev >= 0 {
				for combo := range literCombos[prev] {
					if lib.HasBit(combo, i) {
						combos[lib.SetBit(combo, i, false)] = struct{}{}
					}
				}
			}
		}
		literCombos[liters] = combos
	}
	fmt.Println(len(literCombos[total]))
}

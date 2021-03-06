package main

import (
	"crypto/md5"
	"fmt"

	"github.com/derat/advent-of-code/lib"
)

func main() {
	const (
		nrow = 4
		ncol = 4
		sr   = 0
		sc   = 0
		tr   = 3
		tc   = 3
	)

	pass := lib.InputLines("2016/17")[0]

	// I don't think that A* will work here, since the availability of different
	// directions at a given position changes depending on the path that we took
	// to get there. Unfortunately, I didn't realize this until after writing A*.
	// Just use BFS instead.
	type state struct {
		r, c int
		path string
	}

	var shortest string
	maxSteps := -1
	todo := []state{{0, 0, ""}}
	for steps := 0; true; steps++ {
		// For part 2, we need to find the longest path that reaches the exit.
		// I don't think that there's any way to eliminate paths, since every
		// unique path has its own unpredictable future open/closed state for
		// doors. Just continue trying iterating until we don't have anything
		// left to try (because all earlier paths either reached the exit or
		// ended up in a dead-end state where all four doors are closed).
		if len(todo) == 0 {
			break
		}
		var nextTodo []state

		for _, st := range todo {
			// Check if we're done.
			if st.r == tr && st.c == tc {
				if shortest == "" {
					shortest = st.path
				}
				maxSteps = lib.Max(maxSteps, len(st.path))
				continue
			}

			try := func(dr, dc int, open bool, dir rune) {
				// Skip if we'd be moving outside the outer walls or if the door isn't open.
				next := state{st.r + dr, st.c + dc, st.path + string(dir)}
				if open && next.r >= 0 && next.c >= 0 && next.r < nrow && next.c < ncol {
					nextTodo = append(nextTodo, next)
				}
			}

			h := md5.Sum([]byte(pass + st.path))
			try(-1, 0, lib.Hi(h[0]) >= 0xb, 'U')
			try(1, 0, lib.Lo(h[0]) >= 0xb, 'D')
			try(0, -1, lib.Hi(h[1]) >= 0xb, 'L')
			try(0, 1, lib.Lo(h[1]) >= 0xb, 'R')
		}

		todo = nextTodo
	}

	fmt.Println(shortest)
	fmt.Println(maxSteps)
}

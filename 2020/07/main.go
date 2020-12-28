package main

import (
	"strings"

	"github.com/derat/advent-of-code/lib"
)

type bagInfo struct {
	color string
	num   int
}

func main() {
	holders := make(map[string][]string)
	bags := make(map[string][]bagInfo)

	for _, ln := range lib.InputLines("2020/7") {
		var outer, lst string
		lib.Parse(ln, `^(.+) bags contain (.+)\.$`, &outer, &lst)
		if lst == "no other bags" {
			continue
		}
		for _, p := range strings.Split(lst, ", ") {
			var cnt int
			var inner string
			lib.Parse(p, `^(\d+) (.+) bags?$`, &cnt, &inner)
			holders[inner] = append(holders[inner], outer)
			bags[outer] = append(bags[outer], bagInfo{color: inner, num: cnt})
		}
	}

	seen := make(map[string]struct{})
	var add func(col string)
	add = func(col string) {
		for _, c := range holders[col] {
			if _, ok := seen[c]; !ok {
				seen[c] = struct{}{}
				add(c)
			}
		}
	}
	add("shiny gold")
	println(len(seen))

	var count func(col string) int
	count = func(col string) int {
		total := 0
		for _, b := range bags[col] {
			total += b.num * (1 + count(b.color))
		}
		return total
	}
	println(count("shiny gold"))
}

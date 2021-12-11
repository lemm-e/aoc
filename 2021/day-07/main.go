package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

var set, list []int

func init() {
	data, _ := os.ReadFile("input.txt")
	set = parse(string(data))
	list = seq(extremas(set))
}

func partA() int {
	var fuel int
	for i, v := range list {
		var sum int
		for _, k := range set {
			sum += dist(v, k)
		}
		if i == 0 || sum < fuel {
			fuel = sum
		}
	}
	return fuel
}

func partB() int {
	var spent int
	for i, v := range list {
		var sum int
		for _, k := range set {
			d := dist(v, k)
			sum += fuel(d)
		}
		if i == 0 || sum < spent {
			spent = sum
		}
	}
	return spent
}

func parse(d string) []int {
	list := strings.Split(d, ",")
	var out []int
	for _, v := range list {
		val, _ := strconv.Atoi(v)
		out = append(out, val)
	}
	return out
}

func seq(a, b int) []int {
	min, max := sort(a, b)
	var out []int
	for i := min; i <= max; i++ {
		out = append(out, i)
	}

	return out
}

func sort(a, b int) (min, max int) {
	min, max = a, b
	if a > b {
		min, max = b, a
	}

	return
}

func extremas(list []int) (min, max int) {
	for i, v := range list {
		if i == 0 || v < min {
			min = v
		}

		if i == 0 || v > max {
			max = v
		}
	}

	return
}

func dist(a, b int) int {
	min, max := sort(a, b)
	return max - min
}

func fuel(dist int) int {
	return dist * (dist + 1) / 2
}

func main() {
	fmt.Printf("Part-A: %d\nPart-B: %d", partA(), partB())
}
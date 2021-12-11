package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type board [5][5]int
type input string
type typeStream []int

var (
	data   input
	stream typeStream
	boards map[int]board
)

func init() {
	c, _ := os.ReadFile("input.txt")
	data = input(c)
	stream, boards = data.parse()
}

func (inp input) parse() (typeStream, map[int]board) {
	re := regexp.MustCompile(`\n\n|(?:\r\n){2}`)
	split := re.Split(fmt.Sprint(inp), -1)
	set := make(map[int]board)
	for i, v := range split[1:] {
		var b board
		nums := strings.Fields(v)
		var index int
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				b[i][j], _ = strconv.Atoi(nums[index])
				index++
			}
		}
		set[i] = b
	}

	var inputStream typeStream
	for _, v := range strings.Split(split[0], ",") {
		a, _ := strconv.Atoi(v)
		inputStream = append(inputStream, a)
	}

	return inputStream, set
}

func (b board) isWin() bool {
	for i := 0; i < 5; i++ {
		var sumr, sumc int
		for j := 0; j < 5; j++ {
			sumr += b[i][j]
			sumc += b[j][i]
		}
		if sumr == -5 || sumc == -5 {
			return true
		}
	}
	return false
}

func (b board) scratch(n int) board {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == n {
				b[i][j] = -1
			}
		}
	}
	return b
}

func (s typeStream) match() (interface{}, interface{}) {
	for _, v := range s {
		for i, k := range boards {
			boards[i] = k.scratch(v)
			if boards[i].isWin() {
				out := boards[i]
				delete(boards, i)
				return out, v
			}
		}
	}
	return nil, nil
}

func (b board) sum() int {
	var sum int
	for _, v := range b {
		for _, k := range v {
			if k != -1 {
				sum += k
			}
		}
	}
	return sum
}

func main() {
	var first, last board
	var numf, numl int
	l := len(boards)
	for i := 1; len(boards) != 0; i++ {
		out, v := stream.match()
		switch i {
		case 1:
			first = out.(board)
			numf = v.(int)
		case l:
			last = out.(board)
			numl = v.(int)
		}
	}
	fmt.Printf("Part-1) %d\nPart-2) %d", first.sum()*numf, last.sum()*numl)
}
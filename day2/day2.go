package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	raw, err := ioutil.ReadFile("day2/input.txt")
	check(err)

	pgmStr := strings.Split(string(raw), ",")
	var pgm []int
	for i:=0; i < len(pgmStr); i++ {
		val, _ := strconv.Atoi(pgmStr[i])
		pgm = append(pgm, val)
	}

	for noun := 0; noun <= 100; noun ++ {
		for verb := 0; verb <= 100; verb ++ {
			var testPgm []int
			for i:=0; i < len(pgm); i++ {
				testPgm = append(testPgm, pgm[i])
			}
			testPgm[1] = noun
			testPgm[2] = verb

			result := runComputer(testPgm)
			if result == 19690720 {
				println(100 * noun + verb)
				noun = 101
				verb = 101
			}
		}
	}

}

func runComputer(pgm []int) int {
	pc := 0
	quit := false
	for {
		switch pgm[pc] {
		case 1:
			pgm[pgm[pc+3]] = pgm[pgm[pc+1]] + pgm[pgm[pc+2]]
			pc += 4
		case 2:
			pgm[pgm[pc+3]] = pgm[pgm[pc+1]] * pgm[pgm[pc+2]]
			pc +=4
		case 99:
			quit = true
		default:
			fmt.Println("Error!!")
			quit = true
		}
		if quit || pc >= len(pgm) {
			break
		}
	}
	return pgm[0]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

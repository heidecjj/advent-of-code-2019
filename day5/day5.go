package main

import (
	"github.com/badgerodon/collections/queue"
	_ "github.com/golang-collections/collections/queue"
)

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
	part2()
}

func part1() {
	stdin := queue.New()
	stdout := queue.New()
	stdin.Enqueue(1)
	runComputer(readPgm("day5/input.txt"), stdin, stdout)
	for stdout.Len() > 0 {
		fmt.Printf("%d\n", stdout.Dequeue())
	}
}

func part2() {
	stdin := queue.New()
	stdout := queue.New()
	stdin.Enqueue(5)
	runComputer(readPgm("day5/input.txt"), stdin, stdout)
	for stdout.Len() > 0 {
		fmt.Printf("%d\n", stdout.Dequeue())
	}
}

func readPgm(filename string) []int {
	raw, err := ioutil.ReadFile(filename)
	check(err)

	pgmStr := strings.Split(string(raw), ",")
	var pgm []int
	for i:=0; i < len(pgmStr); i++ {
		val, _ := strconv.Atoi(pgmStr[i])
		pgm = append(pgm, val)
	}
	return pgm
}

func runComputer(pgm []int, stdin, stdout *queue.Queue) []int {
	pc := 0
	quit := false
	for {
		opcode := pgm[pc] % 100
		addrMode := pgm[pc] / 100
		switch opcode {
		case 1:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			pgm[pgm[pc+3]] = args[0] + args[1]
			pc += 4
		case 2:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			pgm[pgm[pc+3]] = args[0] * args[1]
			pc +=4
		case 3:
			pgm[pgm[pc+1]] = stdin.Dequeue().(int)
			pc += 2
		case 4:
			args := handleAddressing(pgm, addrMode, []int{pc+1})
			stdout.Enqueue(args[0])
			pc += 2
		case 5:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			if args[0] != 0 {
				pc = args[1]
			} else {
				pc += 3
			}
		case 6:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			if args[0] == 0 {
				pc = args[1]
			} else {
				pc += 3
			}
		case 7:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			if args[0] < args[1] {
				pgm[pgm[pc+3]] = 1
			} else {
				pgm[pgm[pc+3]] = 0
			}
			pc += 4
		case 8:
			args := handleAddressing(pgm, addrMode, []int{pc+1, pc+2})
			if args[0] == args[1] {
				pgm[pgm[pc+3]] = 1
			} else {
				pgm[pgm[pc+3]] = 0
			}
			pc += 4
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
	return pgm
}

func handleAddressing(pgm []int, addrMode int, argAddrs []int) []int {
	var realArgs []int
	for i := 0; i < len(argAddrs); i++ {
		var realArg int
		if addrMode % 10 == 0 {
			realArg = pgm[pgm[argAddrs[i]]]
		} else {
			realArg = pgm[argAddrs[i]]
		}
		realArgs = append(realArgs, realArg)
		addrMode /= 10
	}
	return realArgs
}

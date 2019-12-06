package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("day1/input.txt")
	check(err)
	defer f.Close()

	fuel := 0
	mass := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err = strconv.Atoi(scanner.Text())
		check(err)
		modFuel := mass / 3 - 2
		addtlFuel := modFuel
		for {
			addtlFuel = addtlFuel/ 3 - 2

			if addtlFuel <= 0 {
				break
			}
			modFuel += addtlFuel
		}

		fuel += modFuel
	}


	fuelStr := strconv.Itoa(fuel)
	fmt.Println(fuelStr)
}

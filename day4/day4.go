package main

import (
	"fmt"
	"strconv"
)

func main() {
	part1(136818, 685879)
	part2(136818, 685879)
}

func part1(start, end int) {
	var passcodes []int
	for candidate := start; candidate <= end; candidate++ {
		if adjacentRule(candidate) && increasingRule(candidate) {
			passcodes = append(passcodes, candidate)
		}
	}
	fmt.Println(strconv.Itoa(len(passcodes)))
}

func part2(start, end int) {
	var passcodes []int
	for candidate := start; candidate <= end; candidate++ {
		if adjacentRule2(candidate) && increasingRule(candidate) {
			passcodes = append(passcodes, candidate)
		}
	}
	fmt.Println(strconv.Itoa(len(passcodes)))
}

func adjacentRule(num int) bool {
	digits := numToDigits(num)
	for i := 0; i < len(digits) - 1; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}
	return false
}

func increasingRule(num int) bool {
	digits := numToDigits(num)
	for i := 0; i < len(digits) - 1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return true
}

func adjacentRule2(num int) bool {
	digits := numToDigits(num)
	lookingAt := digits[0]
	count := 1
	for i := 1; i < len(digits); i++ {
		if digits[i] == lookingAt {
			count++
		} else {
			if count == 2 {
				return true
			}
			lookingAt = digits[i]
			count = 1
		}
	}

	return count == 2
}

func numToDigits(num int) []int {
	var digits []int
	for {
		if num == 0 {
			return digits
		}
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}
}
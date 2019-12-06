package main

import (
	"errors"
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
	both()
}

type Point struct {
	x int
	y int
	steps int
}

func (point Point) cross(point2 Point) int {
	return point.x * point2.y - point.y * point2.x
}
func (point Point) manhattan() int {
	return Abs(point.x) + Abs(point.y)
}
func (point Point) dist(point2 Point) int {
	return Abs(point.x - point2.x) + Abs(point.y - point2.y)
}
func (point Point) print() {
	fmt.Printf("%v\n", point)
}

type Line struct {
	s Point
	e Point
}

func (line Line) vect() Point {
	return Point{line.e.x - line.s.x, line.e.y - line.s.y, line.s.steps}
}


func both() {
	input, err := ioutil.ReadFile("day3/input.txt")
	check(err)

	wireStrs := strings.Split(string(input), "\n")
	var wires [][]string
	var pointss [][]Point
	for i := 0; i < len(wireStrs); i++ {
		wires = append(wires, strings.Split(wireStrs[i], ","))
		pointss = append(pointss, genPoints(wires[i]))
	}
	intersections := findIntersections(pointss[0], pointss[1])
	part1(intersections)
	part2(intersections)
}

func part1(intersections []Point) {
	minPoint := intersections[0]

	for i := 0; i < len(intersections); i++ {
		if intersections[i].manhattan() < minPoint.manhattan() {
			minPoint = intersections[i]
		}
	}
	fmt.Print("Part1 point: ")
	minPoint.print()
	fmt.Printf("Part1 ans  : %d\n", minPoint.manhattan())
}

func part2(intersections []Point) {
	minPoint := intersections[0]

	for i := 0; i < len(intersections); i++ {
		if intersections[i].steps < minPoint.steps {
			minPoint = intersections[i]
		}
	}
	fmt.Print("Part2 point: ")
	minPoint.print()
	fmt.Printf("Part2 ans  : %d\n", minPoint.steps)
}

func genPoints(moves []string) []Point {
	var points []Point
	origin := Point{0, 0 ,0}
	points = append(points, origin)
	for i := 0; i < len(moves); i++ {
		move := moves[i]
		direction := string(move[0])
		magnitude, err := strconv.Atoi(string(move[1:]))
		check(err)
		var newPoint Point
		newPoint.steps = points[i].steps + magnitude
		switch direction {
		case "U":
			newPoint.x = points[i].x
			newPoint.y = points[i].y + magnitude
		case "D":
			newPoint.x = points[i].x
			newPoint.y = points[i].y - magnitude
		case "L":
			newPoint.x = points[i].x - magnitude
			newPoint.y = points[i].y
		case "R":
			newPoint.x = points[i].x + magnitude
			newPoint.y = points[i].y
		}
		points = append(points, newPoint)
	}
	return points
}

func findIntersections(points1, points2 []Point) []Point {
	var intersections []Point
	for i := 0; i < len(points1) - 1; i++ {
		line := Line{points1[i], points1[i+1]}
		for j := 0; j < len(points2) - 1; j++ {
			intersection, err := line.intersection(Line{points2[j], points2[j+1]})
			if err == nil {
				intersections = append(intersections, intersection)
			}
		}
	}

	return intersections
}

func (line Line) intersection(line2 Line) (Point, error) {
	vect1 := line.vect()
	vect2 := line2.vect()
	if vect1.cross(vect2) == 0 {
		return Point{0, 0, 0}, errors.New("lines are parallel")
	}
	var horizontal Line
	var vertical Line
	if line.s.x == line.e.x {
		horizontal = line2
		vertical = line
	} else {
		horizontal = line
		vertical = line2
	}

	if inBetween(horizontal.s.x, horizontal.e.x, vertical.s.x) && inBetween(vertical.s.y, vertical.e.y, horizontal.s.y) {
		intersection := Point{vertical.s.x, horizontal.s.y, 0}
		intersection.steps = horizontal.s.steps + vertical.s.steps + intersection.dist(vertical.s) + intersection.dist(horizontal.s)
		return intersection, nil
	}
	return Point{0, 0, 0}, errors.New("no intersection")
}

func inBetween(start, end, unsure int) bool {
	return start <= unsure && unsure <= end || end <= unsure && unsure <= start
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

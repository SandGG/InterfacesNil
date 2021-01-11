package main

import (
	"fmt"
	"reflect"
)

type people interface {
	average()
}

type student struct {
	name                  string
	cal1, cal2, cal3, avg int
}

func (s *student) average() {
	s.avg = (s.cal1 + s.cal2 + s.cal3) / 3
}

func printPeople(p people) {
	if !(p == nil || reflect.ValueOf(p).IsNil()) {
		p.average()
	}
	fmt.Printf("(%v, %T)\n", p, p)
}

func main() {
	var t *student
	var p people = t
	printPeople(&student{name: "Ana Gomez", cal1: 9, cal2: 10, cal3: 10})
	printPeople(p)
}

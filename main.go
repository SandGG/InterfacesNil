package main

import "fmt"

type people interface {
	average()
}

type student struct {
	name                  string
	cal1, cal2, cal3, avg int
}

func (s *student) average() {
	if s != nil {
		*&s.avg = (s.cal1 + s.cal2 + s.cal3) / 3
	}
}

func printPeople(p people, i inter) {
	if i == nil {
		people.average(p)
		fmt.Printf("(%v, %T)\n", p, p)
	}
}

type inter interface {
	method()
}

func describe(i inter) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var t *student
	var p people = t
	var i inter

	printPeople(&student{name: "Ana Gomez", cal1: 9, cal2: 10, cal3: 10}, i)

	printPeople(p, i)
	describe(i)
	i.method()

}

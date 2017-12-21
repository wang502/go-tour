package main

import (
	"fmt"
)

type bsInterface interface {
	Len() int
	Less(mid int, target int) bool
	Equal(mid int, target int) bool
	Get(i int) interface{}
}

func binarySearch(data bsInterface, target int) interface{} {
	l, r := 0, data.Len()-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if data.Equal(mid, target) {
			return data.Get(mid)
		} else if data.Less(mid, target) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nil
}

type employee struct {
	name string
	age  int
}

func (e *employee) String() string {
	return fmt.Sprintf("%s: %d", e.name, e.age)
}

type employeeAgeFinder struct {
	employees []employee
}

func (f *employeeAgeFinder) Len() int {
	return len(f.employees)
}

func (f *employeeAgeFinder) Less(mid int, target int) bool {
	return f.employees[mid].age < target
}

func (f *employeeAgeFinder) Equal(mid int, target int) bool {
	return f.employees[mid].age == target
}

func (f *employeeAgeFinder) Get(i int) interface{} {
	return f.employees[i]
}

var employees = []employee{
	{"Natalie", 21},
	{"Seth", 24},
	{"Claire", 24},
	{"Maria", 30},
	{"jane", 35},
	{"joan", 40},
}

var finder = &employeeAgeFinder{employees}

func main() {
	res := binarySearch(finder, 40)
	e, _ := res.(employee)
	fmt.Println(e)
}

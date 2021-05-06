package main

import (
	"fmt"
	"reflect"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var employeeList = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}
	return sum
}

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	//["HAO", "CHEN", "MEGAEASE"]

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
	//[3, 4, 8]

	var list1 = []string{"Hao", "Chen", "MegaEase"}
	z := Reduce(list1, func(s string) int { return len(s) })
	fmt.Printf("%v\n", z) // 15

	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool { return n%2 == 1 })
	fmt.Printf("%v\n", out)
	out = Filter(intset, func(n int) bool { return n > 5 })
	fmt.Printf("%v\n", out)

	old := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)
	//old people: 2

	high_pay := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Salary > 6000
	})
	fmt.Printf("High Salary people: %d\n", high_pay)
	//High Salary people: 4

	no_vacation := EmployeeFilterIn(employeeList, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", no_vacation)
	//People no vacation: [{Hao 44 0 8000} {Jack 26 0 4000} {Marry 29 0 6000}]

	total_pay := EmployeeSumIf(employeeList, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", total_pay)
	//Total Salary: 43500

	younger_pay := EmployeeSumIf(employeeList, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("younger_pay: %d\n", younger_pay)

	// Map
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}

	squared_arr := Map(nums, square)
	fmt.Println(squared_arr)
	//[1 4 9 16]

	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
	//[HAO CHEN MEGAEASE]
}

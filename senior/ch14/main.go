package main

import "fmt"

type person struct {
	name string
	age int
}

func NewPerson(name string,age int) *person{
	p:=new(person)
	p.name = name
	p.age = age
	return p
}
func main()  {
	var s string
	s = "张三"
	fmt.Println(s)
	fmt.Printf("%p\n",&s)

	var sp *string
	//*sp = "飞雪无情"
	//fmt.Println(*sp)

	sp = new(string)//关键点
	*sp = "飞雪无情"
	fmt.Println(*sp)

	pp:=NewPerson("飞雪无情",20)
	fmt.Println(pp)

	m:=make(map[string]int,10)
	m["start"] = 1

	fmt.Println(m["start"])
}

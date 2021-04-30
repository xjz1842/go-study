package main

import "fmt"

type address struct {
	province string
	city string
}
func (addr address) String()  string{
	return fmt.Sprintf("the addr is %s%s",addr.province,addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func modifyPerson(p person)  {
	p.name = "李四"
	p.age = 20
}

func modifyPerson1(p *person)  {
	p.name = "李四"
	p.age = 20
}

type person struct {
	name string
	age int
}

func modifyPerson2(p person)  {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n",&p)
	p.name = "李四"
	p.age = 20
}

func modifyPerson3(p *person)  {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n",p)
	p.name = "李四"
	p.age = 20
}

func modifyMap(p map[string]int)  {
	p["飞雪无情"] =20
}

func main()  {
	name:="飞雪无情"
	nameP:=&name//取地址
	fmt.Println("name变量的值为:",name)
	fmt.Println("name变量的内存地址为:",nameP)

	add := address{province: "北京", city: "北京"}
	printString(add)
	printString(&add)

	var si fmt.Stringer = address{province: "上海",city: "上海"}
	printString(si)
	//sip:=&si
	//printString(sip)

	p:=person{name: "张三",age: 18}
	//modifyPerson(p)
	modifyPerson1(&p)
	fmt.Println("person name:",p.name,",age:",p.age)

	p1:=person{name: "张三",age: 18}
	fmt.Printf("main函数：p1的内存地址为%p\n",&p1)
	modifyPerson2(p)
	fmt.Println("person name:",p1.name,",age:",p1.age)

	p2:=person{name: "张三",age: 18}
	fmt.Printf("main函数：p2的内存地址为%p\n",&p2)
	modifyPerson3(&p2)
	fmt.Println("person name:",p2.name,",age:",p2.age)

	m:=make(map[string]int)
	m["飞雪无情"] = 18
	fmt.Println("飞雪无情的年龄为",m["飞雪无情"])
	modifyMap(m)
	fmt.Println("飞雪无情的年龄为",m["飞雪无情"])
}

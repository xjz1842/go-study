package main

import "fmt"

type person struct {
	name string
	age  uint
	address
}
type address struct {
	province string
	city     string
}
type Stringer interface {
	String() string
}

//func (p person) String() string {
//	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
//}
//
//func (addr address) String() string {
//	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
//}
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}
func NewPerson(name string) *person {
	return &person{name: name}
}

func (p *person) String()  string{
	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
}
func (addr address) String()  string{
	return fmt.Sprintf("the addr is %s%s",addr.province,addr.city)
}

type WalkRun interface {
	Walk()
	Run()
}
func (p *person) Walk(){
	fmt.Printf("%s能走\n",p.name)
}
func (p *person) Run(){
	fmt.Printf("%s能跑\n",p.name)
}
func main() {
	p := person{
		age:  30,
		name: "飞雪无情",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(p.address.province)

	printString(&p)
	printString(&p)
	printString(p.address)
	//输出：the addr is 北京北京

	p1 := NewPerson("张三")
	printString(p1)
	//像使用自己的字段一样，直接使用
	fmt.Println(p.province)

	var s fmt.Stringer
	s = p1
	p2:=s.(*person)
	fmt.Println(p2)

	a,ok:=s.(address)
	if ok {
		fmt.Println(a)
	}else {
		fmt.Println("s不是一个address")
	}
}

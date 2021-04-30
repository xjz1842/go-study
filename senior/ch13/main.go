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

func main()  {
	add := address{province: "北京", city: "北京"}
	printString(add)
	printString(&add)
}
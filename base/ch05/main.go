package main

import (
	"errors"
	"fmt"
)

func sum1(a int,b int) int{
	return a+b
}

func sum(a, b int) (int,error){
	if a<0 || b<0 {
		return 0,errors.New("a或者b不能是负数")
	}
	return a + b,nil
}

func sum2(a, b int) (sum int,err error){
	if a<0 || b<0 {
		return 0,errors.New("a或者b不能是负数")
	}
	sum=a+b
	err=nil
	return
}

func sum3(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func colsure() func() int {
	i:=0
	return func() int {
		i++
		return i
	}
}

type Age uint
func (age Age) String(){
	fmt.Println("the age is",age)
}

func (age Age) Update(){
	age = Age(30)
}

func (age *Age) Modify(){
	*age = Age(30)
}


func main()  {
	result:=sum1(1,2)
	fmt.Println(result)

	result,err := sum(1, 2)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("飞雪")
	fmt.Println("飞雪","无情")

	fmt.Println(sum3(1,2))
	fmt.Println(sum3(1,2,3))
	fmt.Println(sum3(1,2,3,4))

	sum5 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum5(1, 2))

	cl:=colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age:=Age(25)
	age.String()
	age.Update()
	age.String()

	//方法赋值给变量，方法表达式
	sm:=Age.String
	//通过变量，要传一个接收者进行调用也就是age
	sm(age)
}

package main

import "fmt"

func main() {
	if i:=6; i >10 {
		fmt.Println("i>10")
	} else if  i>5 && i<=10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}

	switch j:=1;j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}

	switch 2>1 {
	case true:
		fmt.Println("2>1")
	case false:
		fmt.Println("2<=1")
	}
	//for循环
	sum:=0
	for i:=1;i<=100;i++ {
		sum+=i
	}
	fmt.Println("the sum is",sum)
	//sum:=0
	//i:=1
	//for i<=100 {
	//	sum+=i
	//	i++
	//}
	//fmt.Println("the sum is",sum)
	//sum:=0
	//i:=1
	//for {
	//	sum+=i
	//	i++
	//	if i>100 {
	//		break
	//	}
	//}
	//fmt.Println("the sum is",sum)

}
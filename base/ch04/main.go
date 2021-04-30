package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	sum := 0
	for i:=1; i<100; i++{
		if i%2!=0 {
			continue
		}
		sum+=i
	}
	fmt.Println("the sum is",sum)

	array:=[5]string{"a","b","c","d","e"}
	fmt.Println(array[2])
	array1:=[5]string{1:"b",3:"d"}
	for i:=0;i<5;i++{
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array1[i])
	}

	for i,v:=range array{
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	// slice
	array2:=[5]string{"a","b","c","d","e"}
	slice:=array2[2:5]
	fmt.Println(slice)

	//slice1:=make([]string,4)
	//slice2:=make([]string,4,8)
	slice1:=[]string{"a","b","c","d","e"}
	fmt.Println(len(slice1),cap(slice1))

	////追加一个元素
	//slice2:=append(slice1,"f")
	////多加多个元素
	//slice2:=append(slice1,"f","g")
	////追加另一个切片
	//slice2:=append(slice1,slice...)
	// map
	nameAgeMap:=map[string]int{"飞雪无情":20}
	//添加键值对或者更新对应 Key 的 Value
	nameAgeMap["飞雪无情"] = 20
	//获取指定 Key 对应的 Value
	age:=nameAgeMap["飞雪无情"]

	nameAgeMap1:=make(map[string]int)
	nameAgeMap1["飞雪无情1"] = 20
	age,ok:=nameAgeMap1["飞雪无情1"]
	if ok {
		fmt.Println(age)
	}
	//删除
	delete(nameAgeMap,"飞雪无情")

	//测试 for range
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22
	for k,v:=range nameAgeMap{
		fmt.Println("Key is",k,",Value is",v)
	}
	fmt.Println(len(nameAgeMap))

	s:="Hello飞雪无情"
	bs:=[]byte(s)
	fmt.Println(bs)
	fmt.Println(s[0],s[1],s[15])
	fmt.Println(utf8.RuneCountInString(s))

	for i,r:=range s{
		fmt.Println(i,r)
	}

	aa:=[3][3]int{}
	aa[0][0] =1
	aa[0][1] =2
	aa[0][2] =3
	aa[1][0] =4
	aa[1][1] =5
	aa[1][2] =6
	aa[2][0] =7
	aa[2][1] =8
	aa[2][2] =9
	fmt.Println(aa)

}

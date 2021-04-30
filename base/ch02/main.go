package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	j int= 0
	k int= 1
)
//浮点型
var f32 float32 = 2.2
var f64 float64 = 10.3456
//bool类型
var bf bool = false
var bt bool = true
//字符串
var s1 string = "Hello"
var s2 string = "世界"
//零值
var zi int
var zf float64
var zb bool
var zs string
//变量
const name =  "飞雪无情"
const(
	one = iota+1
	two
	three
	four
)

func main() {
	//类型推导
	var i  = 10
	fmt.Println(i)

	fmt.Println(j,k)

	fmt.Println("f32 is",f32,",f64 is",f64)

	fmt.Println("bf is",bf,",bt is",bt)

	fmt.Println("s1 is",s1,",s2 is",s2)
	fmt.Println("s1+s2=",s1+s2)

	fmt.Println(zi,zf,zb,zs)

	fmt.Println(one,two,three,four)
	//转换
	i2s:=strconv.Itoa(i)
	s2i,err:=strconv.Atoi(i2s)
	fmt.Println(i2s,s2i,err)
	i2f:=float64(i)
	f2i:=int(f64)
	fmt.Println(i2f,f2i)

	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1,"H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1,"o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
	//把s1是否包含
	fmt.Println(strings.Index(s1,"H"))
}
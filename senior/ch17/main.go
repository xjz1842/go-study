package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}

func arrayF(a [2]string){
	fmt.Printf("函数arrayF数组指针：%p\n",&a)
}
func sliceF(s []string){
	fmt.Printf("函数sliceF Data：%d\n",(*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

func main()  {

	a2:=[2]string{"飞雪无情","张三"}
	s1:=a2[0:1]
	s2:=a2[:]
	//打印出s1和s2的Data值，是一样的
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	sh1:=(*slice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.Data,sh1.Len,sh1.Cap)

	a1:=[2]string{"飞雪无情","张三"}
	fmt.Printf("函数main数组指针：%p\n",&a1)
	arrayF(a1)
	s3:=a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s3)).Data)
	sliceF(s3)

	s:="飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b:=[]byte(s)
	fmt.Printf("b的内存地址：%d\n",(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s4:=string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)

	s5:="飞雪无情"
	b4:=[]byte(s5)
	//s3:=string(b)
	s6:=*(*string)(unsafe.Pointer(&b4))
	fmt.Println(s6)

	s7:="飞雪无情"
	//b:=[]byte(s)
	sh:=(*reflect.SliceHeader)(unsafe.Pointer(&s7))
	sh.Cap = sh.Len
	b1:=*(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)
	fmt.Println(sh)
}

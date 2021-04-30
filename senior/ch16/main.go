package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type person struct {
	Name string
	Age int
}

func (p person) Print(prefix string){
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}

type Human struct {
	sex  bool
	age  uint8
	min  int
	name string
}

func main()  {
	p:=person{Name: "飞雪无情",Age: 20}
	pv:=reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint:=pv.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)

	i:= 10
	ip:=&i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)

	p1 :=new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName:=(*string)(unsafe.Pointer(p1))
	*pName="飞雪无情"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p1))+unsafe.Offsetof(p1.Age)))
	*pAge = 20
	fmt.Println(*p1)

	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(string("飞雪无情")))
	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情","张三"}))

	h := Human{
		true,
		30,
		1,
		"hello",
	}
	ii := unsafe.Sizeof(h)
	fmt.Println("=====" ,unsafe.Sizeof(h.name))
	jj := unsafe.Alignof(h.age)
	kk := unsafe.Offsetof(h.name)
	fmt.Println(ii, jj, kk)
	fmt.Printf("%p\n", &h)
	var pp unsafe.Pointer
	pp = unsafe.Pointer(&h)
	fmt.Println(pp)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type person struct {
	Name string `json:"name" bson:"b_name"`
	Age int `json:"age" bson:"b_name"`
}

func (p person) String() string{
	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
}


func main()  {
	i:=3
	iv:=reflect.ValueOf(i)
	it:=reflect.TypeOf(i)
	fmt.Println(iv,it)//3 int

	i1:=3
	ipv:=reflect.ValueOf(&i1)
	ipv.Elem().SetInt(4)
	fmt.Println(i1)

	p:=person{Name: "飞雪无情",Age: 20}
	ppv:=reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("张三")
	fmt.Println(p)

	pt:=reflect.TypeOf(p)
	//遍历person的字段
	for i:=0;i<pt.NumField();i++{
		fmt.Println("字段：",pt.Field(i).Name)
	}
	//遍历person的方法
	for i:=0;i<pt.NumMethod();i++{
		fmt.Println("方法：",pt.Method(i).Name)
	}

	pt1:=reflect.TypeOf(p)
	stringerType:=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType:=reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：",pt1.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：",pt1.Implements(writerType))

	//struct to json
	jsonB,err:=json.Marshal(p)
	if err==nil {
		fmt.Println(string(jsonB))
	}
	//json to struct
	respJSON:="{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON),&p)
	fmt.Println(p)

	//遍历person字段中key为json的tag
	for i:=0;i<pt.NumField();i++{
		sf:=pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n",sf.Name,sf.Tag.Get("json"))
	}

	//遍历person字段中key为json、bson的tag
	for i:=0;i<pt.NumField();i++{
		sf:=pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n",sf.Name,sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n",sf.Name,sf.Tag.Get("bson"))
	}

	// struct to json
	pv1:=reflect.ValueOf(p)
	pt2:=reflect.TypeOf(p)
	//自己实现的struct to json
	jsonBuilder:=strings.Builder{}
	jsonBuilder.WriteString("{")
	num:=pt2.NumField()
	for i:=0;i<num;i++{
		jsonTag:=pt2.Field(i).Tag.Get("json") //获取json tag
		jsonBuilder.WriteString("\""+jsonTag+"\"")
		jsonBuilder.WriteString(":")
		//获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"",pv1.Field(i)))
		if i<num-1{
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())//打印json字符串
}
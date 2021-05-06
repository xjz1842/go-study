package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type VisitorFunc func(*Info, error) error

type Visitor1 interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor Visitor1
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor struct {
	visitor Visitor1
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor1
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

type DecoratedVisitor struct {
	visitor    Visitor1
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor1, fn ...VisitorFunc) Visitor1 {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}

	info := Info{}
	var v Visitor1 = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)

	info1 := Info{}
	var v1 Visitor1 = &info1
	v1 = NewDecoratedVisitor(v)
	v1.Visit(loadFile)
}

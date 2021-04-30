package main

import "os"

const name = "飞雪无情"

func main()  {
	os.Mkdir("tmp",0666)
}

func newString() *string{
	s:=new(string)
	*s = "飞雪无情"
	return s
}

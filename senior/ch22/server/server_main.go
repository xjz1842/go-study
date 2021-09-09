package main

import (
	"go-study/senior/ch22"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main()  {
	rpc.RegisterName("MathService",new(mathService.MathService))
	rpc.HandleHTTP()//新增的
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//rpc.Accept(l)
	http.Serve(l, nil)//换成http的服务
}
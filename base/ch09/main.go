package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源
//var sum = 0

var(
	sum int
	mutex sync.Mutex
)

//增加了一个读取sum的函数，便于演示并发
func readSum() int {
	mutex.Lock()
	defer mutex.Unlock()
	b:=sum
	return b
}

func run(){
	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i:=0; i<10;i++ {
		go fmt.Println("和为:",readSum())
	}
	time.Sleep(2 * time.Second)
}

func runCh(){
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}
	for i:=0; i<10;i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为:",readSum())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

//10个人赛跑，1个裁判发号施令
func race(){
	cond :=sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i:=0;i<10; i++ {
		go func(num int) {
			defer  wg.Done()
			fmt.Println(num,"号已经就位")
			cond.L.Lock()
			cond.Wait()//等待发令枪响
			fmt.Println(num,"号开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2*time.Second)
	go func() {
		defer  wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast()//发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}


func main() {
	//开启100个协程让sum+10
	for i := 0; i < 100; i++ {
		go add(10)
	}
	//防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为:",sum)

	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i:=0; i<100;i++ {
		go fmt.Println("和为:",readSum())
	}
	time.Sleep(2 * time.Second)

	runCh()
	doOnce()

	testMap := sync.Map{}
	testMap.Store("test","value")
	fmt.Println(testMap.Load("test"))
}
func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}
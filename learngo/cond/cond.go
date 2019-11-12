package main

import (
	"fmt"
	"sync"
)

var (
	CDT  = false
	wg   = sync.WaitGroup{}
	cond = sync.NewCond(&sync.Mutex{})
)

func main() {
	wg.Add(3)

	go func() {
		defer wg.Done()
		B()
	}()

	go func() {
		defer wg.Done()
		A()
	}()

	wg.Wait()
	fmt.Println("main() end...")
}

func A() {
	cond.L.Lock()
	for !CDT {
		fmt.Println("A() 检测到条件不满足, 等待中...")
		cond.Wait()

		fmt.Println("A() 条件满足, 退出 for...")
	}
	cond.L.Unlock()
}

func B() {
	cond.L.Lock()
	CDT = true
	cond.Signal()
	fmt.Println("B() 通知 A() 条件已满足...")
	cond.L.Unlock()
}

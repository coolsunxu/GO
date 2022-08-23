package main
 
import (
	"log"
	"sync"
	"time"
)
 
func main() {
 
	// 等待组
	var wg sync.WaitGroup
	// 设置超时时间为4
	timeout := 4
	// 等待组中添加两个任务
	wg.Add(2)
	// 新建一个int型通道，大小为1
	done := make(chan int, 1)
 
	// 协程1
	go func() {
		// 函数执行完调用，wg中的计数器减1
		defer wg.Done()
		log.Println("hello")
	}()
 
	// 协程2
	go func() {
		// 函数执行完调用，wg中的计数器减1
		defer wg.Done()
		// 循环操作，直到超时
		times := 1
		for {
			times++
			log.Printf("times: %d", times)
			time.Sleep(1 * time.Second)
		}
	}()
 
	// 协程3
	go func() {
		// 一直阻塞，直到wg中的计数器为0
		wg.Wait()
		// 向通道中写入数据1
		done <- 1
	}()
 
	select {
	// 当能从通道中读取数据时，表明协程1、2已成功执行完
	case <-done:
		log.Println("sucess done")
	// 如果达到超时时间，打印超时
	case <-time.After(time.Second * time.Duration(timeout)):
		log.Println("timeout")
	}
}

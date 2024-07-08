package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type deletionApprovalItems struct {
	approvalID int
	articleID  int
}

var deletionPools = make(chan deletionApprovalItems, 10)
var wg sync.WaitGroup

func auditAsync(valuesRO <-chan int) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("start auditAsync")
	value := <-valuesRO
	deletionPools <- deletionApprovalItems{
		approvalID: value,
		articleID:  value,
	}
}
func deleteAsync() {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	v := <-deletionPools
	fmt.Println(v)
}
func main() {

	fmt.Println("start superDeletion")
	values := make(chan int, 30)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var value int
			fmt.Scan(&value)
			values <- value
			go auditAsync(values)
		}
	}()

	fmt.Println("operate deletion pool articles")
	wg.Add(1)
	go deleteAsync()
	wg.Add(1)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Printf("当前的协程数量:%d\n", runtime.NumGoroutine())
		}
	}()
	wg.Wait()

}

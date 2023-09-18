package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalcpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalcpu)
	//ganti jumlah thread
	runtime.GOMAXPROCS(20)
	//get toto
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalgo := runtime.NumGoroutine()
	fmt.Println("TotalGoroutine", totalgo)
	group.Wait()

}

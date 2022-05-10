package main

import (
	"fmt"
	"runtime"
	"time"
)

func Process(seconds int) {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(seconds))
	close(done)
}

func main() {
	var seconds string

	fmt.Printf("Digite el tiempo en segundos")
	fmt.Scanf("%s", &seconds)

	fmt.Printf("Time in seconds: %s", seconds)
}

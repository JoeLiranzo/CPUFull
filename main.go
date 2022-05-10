package main

import (
	"fmt"
	"runtime"
	"time"
)

func Process(seconds int8) {
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
	var seconds int8

	fmt.Print("Digite el tiempo en segundos: ")
	fmt.Scanf("%d", &seconds)

	fmt.Println("Processing...")
	Process(seconds)

	fmt.Printf("Processed in : %d seconds", seconds)
}

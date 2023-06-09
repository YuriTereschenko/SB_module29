package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func naturNumSquare(wg *sync.WaitGroup, in chan os.Signal) {
	defer wg.Done()
	i := 1
calc:
	for {
		select {
		case <-in:
			break calc
		default:
			fmt.Printf("%v^2= %v\n", i, i*i)
			i++
		}
	}

}

func main() {
	var wg sync.WaitGroup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	wg.Add(1)
	go naturNumSquare(&wg, sigChan)
	wg.Wait()
	fmt.Println("Выхожу из программы")
}

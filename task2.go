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
	var stop os.Signal
	i := 1
	go func() {
		stop = <-in
	}()

	for stop != syscall.SIGINT {
		fmt.Printf("%v^2= %v\n", i, i*i)
		i++

	}

}

func main() {
	var wg sync.WaitGroup
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT)
	wg.Add(1)
	go naturNumSquare(&wg, sigChan)
	wg.Wait()
	fmt.Println("Выхожу из программы")
}

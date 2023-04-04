package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		fmt.Println("выхожу из программы")
		os.Exit(1)
	}()
	i := 1
	for {
		fmt.Printf("%v^2 = %v\n", i, i*i)
		i++
	}
}

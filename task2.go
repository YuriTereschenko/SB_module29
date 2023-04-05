package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := false
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func(flag *bool) {
		<-sigChan
		*flag = true
	}(&stop)
	i := 1

	for !stop {
		fmt.Printf("%v^2 = %v\n", i, i*i)
		i++
	}
	fmt.Println("выхожу из программы")
}

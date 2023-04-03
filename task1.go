package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	in := scaner(&wg)
	first := square(in, &wg)
	second := product(first, &wg)
	receiver(second, &wg)

	wg.Wait()
}

func scaner(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		var scanStr string
		for {
			_, err := fmt.Scan(&scanStr)
			if err != nil {
				log.Println(err)
				continue
			}
			digit, err := strconv.Atoi(scanStr)
			if err != nil {
				if scanStr == "stop" {
					break
				}
				log.Println(err)
				continue
			}
			fmt.Println("User input: ", digit)
			out <- digit
		}
	}()
	return out
}

func square(in chan int, wg *sync.WaitGroup) chan int {
	wg.Add(1)
	out := make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)
		for value := range in {
			result := value * value
			fmt.Println("Square is: ", result)
			out <- result
		}
	}()
	return out
}

func product(in chan int, wg *sync.WaitGroup) chan int {
	wg.Add(1)
	out := make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)

		for value := range in {
			result := value * 2
			fmt.Println("Product is ", result)
			out <- result
		}
	}()
	return out
}

func receiver(in chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Receiver завершил работу")
			wg.Done()
		}()
		for value := range in {
			fmt.Println("Получатель получил: ", value)
		}
	}()
}

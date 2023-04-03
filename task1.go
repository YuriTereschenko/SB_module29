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
	second := mul(first, &wg)
	receiver(second, &wg)

	wg.Wait()
}

func scaner(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("Завершили работу scaner")
		}()
		defer func() {
			close(out)
			fmt.Println("Закрыли канал scaner")
		}()
		var scan string
		var digit int
		for {
			_, err := fmt.Scan(&scan)
			if err != nil {
				log.Println(err)
				continue
			}
			digit, err = strconv.Atoi(scan)
			if err != nil {
				if scan == "stop" {
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
		defer func() {
			wg.Done()
			fmt.Println("Завершили работу square")
		}()
		defer func() {
			close(out)
			fmt.Println("Закрыли канал square")
		}()
		for value := range in {
			result := value * value
			fmt.Println("Отработала функция square in: ", value, "out: ", result)
			out <- result
		}
	}()
	return out
}

func mul(in chan int, wg *sync.WaitGroup) chan int {
	wg.Add(1)
	out := make(chan int)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("Завершили работу mul")
		}()
		defer func() {
			close(out)
			fmt.Println("Закрыли канал mul")
		}()
		for value := range in {
			result := value * 2
			fmt.Println("Отработала функция mul in: ", value, "out: ", result)
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

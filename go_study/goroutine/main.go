package main

import (
	"fmt"
)

//var POOL = 100
//
//func groutine1(p chan int, group *sync.WaitGroup) {
//	defer group.Done()
//	for i := 1; i <= POOL; i++ {
//		p <- i
//		if i%2 == 1 {
//			fmt.Println("groutine-1:", i)
//		}
//	}
//}
//
//func groutine2(p chan int, group *sync.WaitGroup) {
//	defer group.Done()
//	for i := 1; i <= POOL; i++ {
//		<-p
//		if i%2 == 0 {
//			fmt.Println("groutine-2:", i)
//		}
//	}
//}
//
//func main() {
//	msg := make(chan int)
//	var s sync.WaitGroup
//	s.Add(2)
//	go groutine1(msg, &s)
//	go groutine2(msg, &s)
//	s.Wait()
//
//}

func main() {
	chA, chB, chC := make(chan string), make(chan string), make(chan string)

	go func() {
		for i := 1; ; {
			select {
			case <-chA:
				fmt.Println("[A]:", i)
				chB <- "B"
				i += 3
			}
		}
	}()

	go func() {
		for i := 2; ; {
			select {
			case <-chB:
				fmt.Println("[B]:", i)
				chC <- "C"
				i += 3
			}
		}
	}()

	go func() {
		for i := 3; ; {
			if i == 3 {
				chA <- "A"
			}
			select {
			case <-chC:
				fmt.Println("[C]:", i)
				chA <- "A"
				i += 3
			}
		}
	}()

	for !false {
		fmt.Print()
	}
}
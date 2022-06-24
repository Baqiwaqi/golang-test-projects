package main

import (
	"fmt"
)

// daisy chain
func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

// func main() {

// 	quit := make(chan bool)
// 	joe := boring("Joe!", quit)
// 	// timeout := time.After(time.Second * 5)
// 	for i := rand.Intn(10); i > 0; i-- {
// 		fmt.Printf("You say: %q\n", <-joe)
// 	}
// 	quit <- true

// 	// 	select {
// 	// 	case s := <-joe:
// 	// 		fmt.Println(s)
// 	// 	case <-timeout:
// 	// 		fmt.Println("You talk too much.")
// 	// 		return
// 	// 	}
// 	// }
// 	// ann := boring("Ann !")
// 	// now the function can independantly execute
// 	// c := fanIn(boring("Joe"), boring("Ann"))
// 	// for i := 0; i < 10; i++ {
// 	// 	// fmt.Printf("You say: %q\n", <-c) // receive from channel;
// 	// 	// fmt.Printf("You say: %q\n", <-joe) // receive from channel;
// 	// 	// fmt.Printf("You say: %q\n", <-ann) // receive from channel;
// 	// 	fmt.Printf("You say: %q\n", <-c) // receive from channel;
// 	// }
// 	// fmt.Println("I'm listening.")
// 	// time.Sleep(2 * time.Second)
// 	// fmt.Println("You're boring; I'm leaving.")
// }

// // make the channel an argument to fanIn
// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for {
// 			select {
// 			case s := <-input1:
// 				c <- s
// 			case s := <-input2:
// 				c <- s
// 			}
// 		}
// 	}()
// 	return c
// 	// go func() {
// 	// 	for {
// 	// 		c <- <-input1
// 	// 	}
// 	// }()
// 	// go func() {
// 	// 	for {
// 	// 		c <- <-input2
// 	// 	}
// 	// }()
// 	// return c
// }

// func boring(msg string, quit <-chan bool) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 			select {
// 			case <-quit:
// 				fmt.Println("Bye!")
// 				return
// 			default:
// 			}

// 			time.Sleep(time.Second)
// 		}
// 	}()
// 	return c
// }

// // func boring(msg string) string { //, c chan string) {
// // 	for i := 0; ; i++ {
// // 		// c <- fmt.Sprintf("%s %d", msg, i+1)
// // 		fmt.Printf("%s %d\n", msg, i+1)
// // 		time.Sleep(time.Second)
// // 	}
// // }

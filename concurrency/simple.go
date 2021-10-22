package main

import "fmt"

func saysHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v, says Hello %v times \n", name, i)
	}
}

func saysHello2(name string, pipe chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v, says Hello %v times \n", name, i)
	}

	pipe <- "done hai ji!!"
}

func saysHello3(name string, pipe chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v, says Hello %v times \n", name, i)
	}

	pipe <- 1
}

func main()  {
	// saysHello("Ritik")
	// saysHello("Soumya")


	// go -> use to create goroutines
	// multiple threads are introduced
	// but without wait for any other thread main thread ends
	// go saysHello("Riitk")
	// go saysHello("Soumya")

	// to wait for other threads use channels
	// wait for those channel input to returned
	channel1 := make(chan string)
	channel2 := make(chan int)

	go saysHello2("Ritik", channel1)
	go saysHello3("Soumya", channel2)

	<- channel1
	<- channel2
}
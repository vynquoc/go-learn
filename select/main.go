package main

import (
	"fmt"
	"time"
)

func main() {
	chanGoogle := make(chan string)
	chanBing := make(chan string)

	go googleSearch(chanGoogle)
	go bingSearch(chanBing)

	select {
	case result := <-chanGoogle:
		fmt.Println(result)
	case result := <-chanBing:
		fmt.Println(result)
	}

	fmt.Println("DONE !!!")
}

func googleSearch(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "found from Google"

}

func bingSearch(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "found from Bing"
}

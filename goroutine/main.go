package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	firstChan := make(chan string)
	secondChan := make(chan string)
	thirdChan := make(chan string)

	f, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		return
	}

	listChan := []chan string{firstChan, secondChan, thirdChan}
	listUrl := []string{"https://youtube.com", "https://google.com", "https://daily.dev"}

	for i := 0; i < len(listUrl); i++ {
		go getData(listUrl[i], listChan[i])
	}

	for i := 0; i < len(listUrl); i++ {
		_, err := f.WriteString(<-listChan[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("DONE!!!")
}

func getData(url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println((err))
		return
	}
	content, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		fmt.Println((err))
		return
	}

	ch <- string(content)
	defer res.Body.Close()
}

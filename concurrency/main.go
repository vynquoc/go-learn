package main

import (
	"encoding/json"
	"fmt"
	"go-concurrency/concurrency/visit"
	"io/ioutil"
	"log"
	"sync"
)

type Task struct {
	Date   string
	Visits []visit.Visit
}

type DailyStat struct {
	Date   string         `json:"date"`
	ByPage map[string]int `json:"byPage"`
}

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	dayStats := make(map[string][]visit.Visit)
	err = json.Unmarshal(data, &dayStats)
	if err != nil {
		log.Fatal(err)
	}
	var w8 sync.WaitGroup
	w8.Add(len(dayStats))
	inputCh := make(chan Task, 10)
	outputCh := make(chan DailyStat, len(dayStats))

	for k := 0; k < len(dayStats); k++ {
		go worker(inputCh, k, outputCh, &w8)
	}

	for date, visits := range dayStats {
		inputCh <- Task{
			Date:   date,
			Visits: visits,
		}
	}
	close(inputCh)
	w8.Wait()
	close(outputCh)
	done := make([]DailyStat, 0, len(dayStats))
	for out := range outputCh {
		done = append(done, out)
	}

	res, err := json.Marshal(done)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("results.json", res, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done !")
}

func worker(in chan Task, workerId int, out chan DailyStat, w8 *sync.WaitGroup) {
	for received := range in {
		m := make(map[string]int)
		for _, v := range received.Visits {
			m[v.Page]++
		}
		out <- DailyStat{
			Date:   received.Date,
			ByPage: m,
		}
		fmt.Printf("[worker %d] finished task \n", workerId)
		log.Println("worker quit")
		w8.Done()
	}
}

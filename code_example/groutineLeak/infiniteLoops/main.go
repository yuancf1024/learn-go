package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(url string, wg *sync.WaitGroup) {
	i := 0
	for {
		if _, err := http.Get(url); err == nil {
			// write to db
			break
		}
		i++
		if i >= 3 {
			break
		}
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go request(fmt.Sprintf("https://127.0.0.1:8080/%d", i), &wg)
	}
	wg.Wait()
}
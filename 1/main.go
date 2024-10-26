package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go resettable(&wg)

	wg.Wait()
}

func resettable(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)

	for {
		hasTimeout(2, ticker)
	}

}

func hasTimeout(duration int, ticker *time.Ticker) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()
	ticks(ticker, ctx)
}

func ticks(ticker *time.Ticker, ctx context.Context) {
	fmt.Println("Reset")
	for {
		select {
		case <-ticker.C:
			fmt.Println("running")
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		}

	}
}

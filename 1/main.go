package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	const SessionTimeout = 2

	wg := sync.WaitGroup{}

	wg.Add(1)
	go AppContainer(&wg, SessionTimeout)

	wg.Wait()
}

func AppContainer(wg *sync.WaitGroup, timeout int) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	sessionLifetime := time.Duration(timeout) * time.Second
	for {
		hasTimeout(sessionLifetime, ticker)
	}

}

func hasTimeout(duration time.Duration, ticker *time.Ticker) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
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

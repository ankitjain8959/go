package concurrency

import "log"

func UnderstandGoroutines() {
	for range 5 {
		log.Println("Hello, Goroutines!")
	}
}

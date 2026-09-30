package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	sites := []string{"site1.com", "site2.com", "site3.com", "site4.com", "site5.com", "site6.com"}
	results := make([]string, len(sites))

	for i, site := range sites {
		wg.Add(1)
		go func(i int, site string) {
			defer wg.Done()
			time.Sleep(200 * time.Millisecond)
			results[i] = site + " - OK"
		}(i, site)
	}

	wg.Wait()

	for _, r := range results {
		fmt.Println(r)
	}
}

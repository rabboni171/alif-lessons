package main

import (
	"fmt"
	"sync"
	"time"
)

func checkSite(name string) string {
	time.Sleep(200 * time.Millisecond)
	return name + " - OK"
}

func main() {
	var wg sync.WaitGroup

	sites := []string{"site1.com", "site2.com", "site3.com", "site4.com", "site5.com", "site6.com"}
	results := make([]string, len(sites))

	start := time.Now()

	for i, site := range sites {
		wg.Add(1)
		go func(i int, site string) {
			defer wg.Done()
			results[i] = checkSite(site)
		}(i, site)
	}

	wg.Wait()
	parallelTime := time.Since(start)

	for _, r := range results {
		fmt.Println(r)
	}

	sequentialTime := time.Duration(len(sites)) * 200 * time.Millisecond
	fmt.Println("Параллельно заняло:", parallelTime)
	fmt.Println("По одному заняло бы:", sequentialTime)
}

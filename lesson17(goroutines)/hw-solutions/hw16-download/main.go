package main

import (
	"fmt"
	"sync"
	"time"
)

func download(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	fmt.Println("Скачали:", name)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	files := []string{"file1.zip", "file2.zip", "file3.zip", "file4.zip", "file5.zip"}

	for _, f := range files {
		wg.Add(1)
		go download(f, &wg)
	}

	wg.Wait()

	fmt.Println("Всего заняло:", time.Since(start))       // ~1 секунда
	fmt.Println("По одному заняло бы:", 5*time.Second) // для сравнения
}

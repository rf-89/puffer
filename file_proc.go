package main

import (
	"sync"
)

func fileProc(files []string, out string, n int) []map[string]string {
	var info []map[string]string

	maxExecutions := mathThreads(n)

	var wg = &sync.WaitGroup{}

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		maxExecutions <- struct{}{}
		go func(i int) {
			info = append(info, mathHash(files[i]))
			<-maxExecutions
			wg.Done()
		}(i)
	}
	wg.Wait()

	return info
}

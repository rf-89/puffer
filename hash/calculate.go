package hash

import (
	"sync"

	"github.com/teapod89/puffer/util/math"
)

func Calculate(files []string, out string, n int) []map[string]string {
	var info []map[string]string

	maxExecutions := math.ThreadSum(n)

	var wg = &sync.WaitGroup{}

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		maxExecutions <- struct{}{}
		go func(i int) {
			info = append(info, math.Hash(files[i]))
			<-maxExecutions
			wg.Done()
		}(i)
	}
	wg.Wait()

	return info
}

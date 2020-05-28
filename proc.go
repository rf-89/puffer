package main

import (
	"fmt"
	"runtime"
	"sync"
)

func fileProc(files []string, out string, num int) []map[string]string {
	var info []map[string]string

	//最大並行数の指定
	maxThread := make(chan struct{}, num)
	//ただし、CPU数よりも小さい数が指定されていた場合、CPUコア数を最大数とするようにする。
	cpus := runtime.NumCPU()
	fmt.Println(cpus)
	if num < cpus {
		maxThread = make(chan struct{}, cpus)
	}

	var procNum = len(files)
	var wg = &sync.WaitGroup{}

	for i := 0; i < procNum; i++ {
		wg.Add(1)
		maxThread <- struct{}{}
		go func(i int) {
			info = append(info, mathHash(files[i]))
			<-maxThread
			wg.Done()
		}(i)
	}
	wg.Wait()

	return info
}

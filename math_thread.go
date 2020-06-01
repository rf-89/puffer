package main

import "runtime"

func mathThreads(n int) chan struct{} {
	//最大並行数の指定
	threads := make(chan struct{}, n)
	//ただし、CPU数よりも小さい数が指定されていた場合、CPUコア数を最大数とするようにする。
	cpus := runtime.NumCPU()
	if n < cpus {
		threads = make(chan struct{}, cpus)
	}
	return threads
}

package math

import (
	"runtime"
)

func ThreadSum(n int) chan struct{} {
	//最大並行数の指定
	threads := make(chan struct{}, n)
	//ただし、CPU数よりも小さい数が指定されていた場合、CPUコア数を最大数とするようにする。
	cpus := runtime.NumCPU()
	//利用可能なコアすべてを使って計算する。
	runtime.GOMAXPROCS(cpus)
	if n < cpus {
		threads = make(chan struct{}, cpus)
	}
	return threads
}

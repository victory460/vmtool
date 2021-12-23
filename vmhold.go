package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var a []string

// TODO 通过flag库，将利用率控制参数放到命令行的参数中
func main() {
	cpuCount, _ := cpu.Counts(true)
	fmt.Printf("cpu count : %v \r\n", cpuCount)
	PrintMemUsage()
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent %v \r\n", percent)
	if percent[0] > 30 {
		fmt.Printf("it is fine")
		return
	}
	grtCount := cpuCount / 3
	fmt.Printf("grtCount : %v", grtCount)
	for i := 0; i < grtCount+1; i++ { // cpu > 30%
		go task()
	}
	select {}
}

func task() {
	for {
		vm, _ := mem.VirtualMemory()
		if vm.UsedPercent < 65 { // mem >
			// TODO 这边可以优化
			a = append(a, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		} else if vm.UsedPercent > 85 {
			runtime.GC() //调用 runtime.GC() 进行手动触发GC进行内存回收
		}
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", BToMB(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", BToMB(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", BToMB(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func BToMB(b uint64) uint64 {
	return b / 1024 / 1024
}

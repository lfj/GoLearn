package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 显示当前内存使用情况
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Before GC: Alloc = %v KB\n", memStats.Alloc/1024)

	// 创建大量对象
	for i := 0; i < 1e6; i++ {
		_ = make([]byte, 1024) // 分配 1MB 内存
	}

	// 手动触发垃圾回收
	runtime.GC()

	// 再次检查内存使用情况
	runtime.ReadMemStats(&memStats)
	fmt.Printf("After GC: Alloc = %v KB\n", memStats.Alloc/1024)
}

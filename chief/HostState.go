package chief

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"xiaowumin-SFM/Struct"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetHostState() (*Struct.HostState, error) { // 获取服务器硬件使用状态
	//var HostStateData []int
	//HostStateData = append(HostStateData, getCPUUsage(), getUsedPercent())
	hoststate := Struct.HostState{
		CpuUse:    getCPUUsage(),
		MemoryUse: getUsedPercent(),
	}

	return &hoststate, nil
}

func getCPUUsage() int { // 获取服务器的处理器占用
	cmd := exec.Command("wmic", "cpu", "get", "loadpercentage")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("获取CPU使用率失败:", err)
	}

	cpuUsageStr := strings.Split(string(output), "\n")[1]
	cpuUsageStr = strings.TrimSpace(cpuUsageStr) // 清理空格和制表符

	// 提取数字部分
	cpuUsageStr = strings.Split(cpuUsageStr, " ")[0]
	cpuUsage, err := strconv.Atoi(cpuUsageStr)
	if err != nil {
		fmt.Println("转换内存大小为整数失败:", err)
		return 0
	}
	//fmt.Println("处理器使用率:", cpuUsage)
	return cpuUsage
}

func getUsedPercent() int { //获取服务器内存占用，使用了gopsutil库 https://github.com/shirou/gopsutil
	v, _ := mem.VirtualMemory()
	usedPercent := v.UsedPercent
	//fmt.Printf("内存使用率: %.2f\n", usedPercent)
	return int(usedPercent)
}

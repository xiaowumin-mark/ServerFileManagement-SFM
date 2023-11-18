package chief

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
	"xiaowumin-SFM/Struct"
)

type DiskUsage struct {
	Name       string
	TotalSpace float64
	UsedSpace  float64
	FreeSpace  float64
}

const (
	DRIVE_FIXED   = 3
	DRIVE_UNKNOWN = 0
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var getDiskFreeSpaceEx = kernel32.NewProc("GetDiskFreeSpaceExW")
var getDriveType = kernel32.NewProc("GetDriveTypeW")

func getDiskUsage() ([]DiskUsage, error) {
	var diskUsages []DiskUsage

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	GetLogicalDrives := kernel32.NewProc("GetLogicalDrives")
	bitmask, _, _ := GetLogicalDrives.Call()

	for i := 0; i < 26; i++ {
		if (bitmask>>i)&1 == 0 {
			continue // Continue if the drive doesn't exist
		}

		drive := string(rune('A'+i)) + ":\\"
		driveType, _, _ := getDriveType.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drive))))
		if driveType != DRIVE_FIXED {
			continue // Continue if the drive is not a fixed drive
		}

		lpFreeBytesAvailable := int64(0)
		lpTotalNumberOfBytes := int64(0)
		lpTotalNumberOfFreeBytes := int64(0)

		_, _, _ = getDiskFreeSpaceEx.Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drive))),
			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)),
		)

		diskUsages = append(diskUsages, DiskUsage{
			Name:       drive,
			TotalSpace: float64(lpTotalNumberOfBytes) / (1024 * 1024 * 1024),
			UsedSpace:  (float64(lpTotalNumberOfBytes) - float64(lpTotalNumberOfFreeBytes)) / (1024 * 1024 * 1024),
			FreeSpace:  float64(lpTotalNumberOfFreeBytes) / (1024 * 1024 * 1024),
		})
	}

	return diskUsages, nil
}

func Config() []byte {
	diskUsages, err := getDiskUsage()
	if err != nil {
		fmt.Println("Failed to retrieve disk usage:", err)

	}

	var DiskAll []string
	var DiskName []string
	var DiskFreeSpace []string
	var DiskUse []string
	for i := 0; i < len(diskUsages); i++ {
		DiskAll = append(DiskAll, strconv.FormatFloat(diskUsages[i].TotalSpace, 'f', 2, 64)+" GB")
		DiskName = append(DiskName, diskUsages[i].Name)
		DiskFreeSpace = append(DiskFreeSpace, strconv.FormatFloat(diskUsages[i].FreeSpace, 'f', 2, 64)+" GB")
		DiskUse = append(DiskUse, strconv.FormatFloat(diskUsages[i].UsedSpace, 'f', 2, 64)+" GB")
	}

	fwqc := Struct.Config{
		Disks: struct {
			All     []string `json:"All"`
			Name    []string `json:"Name"`
			Numbers int      `json:"Numbers"`
			Residue []string `json:"Residue"`
			Use     []string `json:"Use"`
		}{
			All:     DiskAll,
			Name:    DiskName,
			Numbers: len(diskUsages),
			Residue: DiskFreeSpace,
			Use:     DiskUse,
		},
	}
	// 打印JSON字符串
	// 将结构体转换为JSON字符串
	jsonData, err := json.Marshal(fwqc)
	if err != nil {
		fmt.Println("转换为JSON时出错:", err)

	}

	// 返回JSON
	return jsonData
}

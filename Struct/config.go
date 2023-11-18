package Struct

type Config struct {
	Disks struct { //硬盘
		All     []string `json:"All"`     //总空间
		Name    []string `json:"Name"`    // 磁盘名称
		Numbers int      `json:"Numbers"` // 磁盘数量
		Residue []string `json:"Residue"` //磁盘剩余空间
		Use     []string `json:"Use"`     // 已使用空间
	} `json:"Disks"`
}

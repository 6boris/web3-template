package net

import (
	"fmt"
	"os/exec"
)

func _portInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}

func GetRandPort(start, end int64) int64 {
	for i := start; i <= end; i++ {
		if !_portInUse(int(i)) {
			return i
		}
	}
	return 0
}

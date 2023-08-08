package utils

import (
	"os"
	"fmt"
	"syscall"
)

func RefreshTeku(pid int) {
	// Find the process with the given PID
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Error finding process:", err)
		return
	}

	// Send SIGHUP signal to the process
	err = process.Signal(syscall.SIGHUP)
	if err != nil {
		fmt.Println("Error sending signal:", err)
		return
	}
}
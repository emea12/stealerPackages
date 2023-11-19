package functions

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CloseBrowser() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		// macOS
		cmd = exec.Command("pkill", "Safari") // change "Safari" to the appropriate process name
	case "windows":
		// Windows
		cmd = exec.Command("taskkill", "/F", "/IM", "chrome.exe") // change "chrome.exe" to the appropriate process name
	case "linux":
		// Linux
		cmd = exec.Command("pkill", "firefox") // change "firefox" to the appropriate process name
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to close browser: %v", err)
	}

	return nil
}

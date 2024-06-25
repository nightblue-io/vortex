package internal

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func IsWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}

	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

func OpenUrl(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		if IsWSL() {
			// Use 'cmd.exe /c start' to open the
			// URL in the default Windows browser.
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			// Use xdg-open on native Linux environments.
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent
		// issues with URLs starting with a quote.
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}

	return exec.Command(cmd, args...).Start()
}

func GetLocalAccessToken() string {
	data, err := os.ReadFile(FileAccessToken)
	if err != nil {
		return ""
	}

	return string(data)
}

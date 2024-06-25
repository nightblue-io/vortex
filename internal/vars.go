package internal

import (
	"os"
	"path/filepath"
)

var (
	homeDir = func() string {
		h, _ := os.UserHomeDir()
		return h
	}

	DirVortex       = filepath.Join(homeDir(), ".vortex")
	FileAccessToken = filepath.Join(DirVortex, "accesstoken")
)

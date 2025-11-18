package Constants

import (
	"os"
	"path/filepath"
)

// GetBaseDir returns the working directory of the running binary
func GetBaseDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

// --- core paths --------------------------------------------------
var BASE_DIR string = GetBaseDir()
var PUBLIC_DIR string = filepath.Join(BASE_DIR, "public")
var UPLOADS_PATH string = filepath.Join(PUBLIC_DIR, "uploads")

// --- application meta -------------------------------------------
var VERSION string = "0.0.5"

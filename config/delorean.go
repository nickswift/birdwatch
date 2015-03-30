package config

import(
	"os"
	"path/filepath"
)

const (
	DELOR_VERSION_MAJOR = 0
	DELOR_VERSION_MINOR = 1
)

var (
	DELOR_DATA_DIR = filepath.Join(os.Getenv("BW_DATA_DIR"), "publ")
)
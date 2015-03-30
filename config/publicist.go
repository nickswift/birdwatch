package config

import(
	"os"
	"path/filepath"
)

const (
	PUBL_VERSION_MAJOR = 0
	PUBL_VERSION_MINOR = 1
)

var (
	PUBL_DATA_DIR = filepath.Join(os.Getenv("BW_DATA_DIR"), "publ")
)
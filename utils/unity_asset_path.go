package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func IsValidAssetPath(path string) bool {
	if path[len(path)-1] == '~' || strings.ToLower(filepath.Ext(path)) == ".tmp" {
		return false
	}

	dirpath, file := filepath.Split(path)

	if file[0] == '.' {
		return false
	}

	if len(dirpath) == 0 {
		return true
	}

	dirs := strings.Split(dirpath, string(os.PathSeparator))

	for _, dir := range dirs {
		if len(dir) == 0 {
			continue
		}
		dir = strings.ToLower(dir)
		if dir[0] == '.' || dir == "cvs" {
			return false
		}
	}

	return true
}

func IsMetaFile(path string) bool {
	return strings.ToLower(filepath.Ext(path)) == ".meta"
}

func RemoveMetaExt(path string) string {
	return path[:len(path)-len(".meta")]
}

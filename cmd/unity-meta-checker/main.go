package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/YuukiARIA/unity-meta-checker/models"
	"github.com/jessevdk/go-flags"
)

func isValidAssetPath(path string) bool {
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

func isMetaFile(path string) bool {
	return strings.ToLower(filepath.Ext(path)) == ".meta"
}

func removeMetaExt(path string) string {
	return path[:len(path)-len(".meta")]
}

func collectAssetFiles(rootPath string) (map[string]models.AssetPathInfo, error) {
	assetPathInfos := make(map[string]models.AssetPathInfo)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if path == rootPath {
			return nil
		}

		relpath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}

		// process an asset file
		assetPath := models.AssetPathInfo{Path: relpath, FileInfo: info}

		if info.IsDir() {
			// found a directory
			assetPath.IsEmpty = true
		} else {
			if isMetaFile(relpath) {
				// found a .meta file
				assetPath.IsMeta = true
			}
		}

		assetPathInfos[relpath] = assetPath
		return nil
	})

	return assetPathInfos, err
}

func main() {
	var opts models.Options
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}

	assetsPath := filepath.Join(opts.ProjectPath, "Assets")

	assetPathInfos, err := collectAssetFiles(assetsPath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", assetPathInfos)
}

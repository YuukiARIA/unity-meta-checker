package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/YuukiARIA/unity-meta-checker/models"
	"github.com/YuukiARIA/unity-meta-checker/utils"
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

		parent := filepath.Dir(relpath)
		parentAssetPathInfo := assetPathInfos[parent]
		parentAssetPathInfo.IsEmpty = false
		assetPathInfos[parent] = parentAssetPathInfo

		// process an asset file
		assetPath := models.AssetPathInfo{
			Path:             relpath,
			IsValidAssetPath: isValidAssetPath(relpath),
			FileInfo:         info,
		}

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

func validateAssetPaths(assetPathInfos map[string]models.AssetPathInfo) ([]string, []string) {
	checkedAssetPathSet := utils.StringSet{}

	danglingMetaPaths := make([]string, 0)
	metalessAssetPaths := make([]string, 0)

	// Enumerate dangling .meta paths
	for path, info := range assetPathInfos {
		if info.IsMeta {
			assetPath := removeMetaExt(path)
			checkedAssetPathSet.Add(assetPath)
			if _, exists := assetPathInfos[assetPath]; !exists {
				danglingMetaPaths = append(danglingMetaPaths, path)
			}
		}
	}

	// Enumerate meta-less asset paths
	for path, info := range assetPathInfos {
		if !info.IsMeta && info.IsValidAssetPath && !checkedAssetPathSet.Contains(path) {
			metalessAssetPaths = append(metalessAssetPaths, path)
		}
	}

	sort.Strings(danglingMetaPaths)
	sort.Strings(metalessAssetPaths)

	return danglingMetaPaths, metalessAssetPaths
}

func main() {
	var opts models.Options
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}

	assetsPath := filepath.Join(opts.ProjectPath, "Assets")
	stat, err := os.Stat(assetsPath)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "ERROR: %s does not exist.\n", assetsPath)
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "ERROR: %s is not directory.\n", assetsPath)
		os.Exit(1)
	}

	assetPathInfos, err := collectAssetFiles(assetsPath)

	if err != nil {
		panic(err)
	}

	danglingMetaPaths, metalessAssetPaths := validateAssetPaths(assetPathInfos)
	fmt.Printf("Dangling Meta Files\n%#v\n", danglingMetaPaths)
	fmt.Printf("Meta-less Assets\n%#v\n", metalessAssetPaths)
}

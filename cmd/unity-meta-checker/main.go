package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/YuukiARIA/unity-meta-checker/models"
	"github.com/YuukiARIA/unity-meta-checker/render"
	"github.com/YuukiARIA/unity-meta-checker/utils"
	"github.com/jessevdk/go-flags"
)

func collectAssetFiles(rootPath string) (map[string]*models.AssetPathInfo, error) {
	assetPathInfos := make(map[string]*models.AssetPathInfo)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if path == rootPath {
			return nil
		}

		relpath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}

		assetPathInfos[relpath] = models.NewAssetPathInfo(relpath, info.IsDir())
		return nil
	})

	return assetPathInfos, err
}

func validateAssetPaths(assetPathInfos map[string]*models.AssetPathInfo) *models.Result {
	checkedAssetPathSet := utils.StringSet{}

	danglingMetaPaths := make([]string, 0)
	metalessAssetPaths := make([]string, 0)

	// Enumerate dangling .meta paths
	for path, info := range assetPathInfos {
		if info.IsMeta {
			assetPath := utils.RemoveMetaExt(path)
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

	return models.NewResult(danglingMetaPaths, metalessAssetPaths)
}

func isHelpFlagGiven(err error) bool {
	flagsErr, ok := err.(*flags.Error)
	return ok && flagsErr.Type == flags.ErrHelp
}

func main() {
	var opts models.Options
	if _, err := flags.Parse(&opts); err != nil {
		if isHelpFlagGiven(err) {
			return
		} else {
			os.Exit(1)
		}
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

	result := validateAssetPaths(assetPathInfos)

	output := os.Stdout
	if len(opts.Output) > 0 {
		if output, err = os.Create(opts.Output); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "Write result to %s\n", opts.Output)
	}

	var t *template.Template
	if opts.TemplatePath != "" {
		t, err = render.LoadTemplate(opts.TemplatePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
	} else {
		t = render.GetDefaultTemplate()
	}
	if err := render.RenderResult(result, t, output); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	if opts.RaiseError && result.HasContent() {
		os.Exit(1)
	}
}

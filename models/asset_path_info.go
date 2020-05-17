package models

import "github.com/YuukiARIA/unity-meta-checker/utils"

type AssetPathInfo struct {
	Path             string
	IsValidAssetPath bool
	IsMeta           bool
}

func NewAssetPathInfo(assetPath string, isDir bool) *AssetPathInfo {
	return &AssetPathInfo{
		Path:             assetPath,
		IsValidAssetPath: utils.IsValidAssetPath(assetPath),
		IsMeta:           !isDir && utils.IsMetaFile(assetPath),
	}
}

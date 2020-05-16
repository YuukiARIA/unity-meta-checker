package models

import "os"

type AssetPathInfo struct {
	Path             string
	IsValidAssetPath bool
	IsMeta           bool
	IsEmpty          bool
	FileInfo         os.FileInfo
}

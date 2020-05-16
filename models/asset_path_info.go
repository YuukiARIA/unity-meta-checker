package models

import "os"

type AssetPathInfo struct {
	Path     string
	IsMeta   bool
	IsEmpty  bool
	FileInfo os.FileInfo
}

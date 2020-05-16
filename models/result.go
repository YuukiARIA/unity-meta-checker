package models

type Result struct {
	DanglingMetaPaths  []string
	MetalessAssetPaths []string
}

func NewResult(danglingMetaPaths, metalessAssetPaths []string) *Result {
	return &Result{danglingMetaPaths, metalessAssetPaths}
}

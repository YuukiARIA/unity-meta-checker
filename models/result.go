package models

type Result struct {
	DanglingMetaPaths  []string
	MetalessAssetPaths []string
}

func NewResult(danglingMetaPaths, metalessAssetPaths []string) *Result {
	return &Result{danglingMetaPaths, metalessAssetPaths}
}

func (r *Result) HasContent() bool {
	return len(r.DanglingMetaPaths) > 0 && len(r.MetalessAssetPaths) > 0
}

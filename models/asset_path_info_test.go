package models

import "testing"

func TestNewAssetPathInfo(t *testing.T) {
	testCases := []struct {
		path                     string
		isDir                    bool
		expectedIsValidAssetPath bool
		expectedIsMeta           bool
	}{
		{"a/b/asset", false, true, false},
		{"a/b/test.meta", false, true, true},

		{"a/b/folder", true, true, false},
		{"a/b/folder.meta", true, true, false},

		{"a/b/xxx~", false, false, false},
		{"a/b/.xyz", true, false, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.path, func(t *testing.T) {
			a := NewAssetPathInfo(testCase.path, testCase.isDir)
			if a.Path != testCase.path {
				t.Errorf("expected .Path is %v, but got %v", testCase.path, a.Path)
			}
			if a.IsValidAssetPath != testCase.expectedIsValidAssetPath {
				t.Errorf("expected .IsValidAssetPath is %v, but got %v", testCase.expectedIsValidAssetPath, a.IsValidAssetPath)
			}
			if a.IsMeta != testCase.expectedIsMeta {
				t.Errorf("expected .IsMeta is %v, but got %v", testCase.expectedIsMeta, a.IsMeta)
			}
		})
	}
}

package utils

import (
	"testing"
)

func TestIsValidAssetPath(t *testing.T) {
	falseExamples := []string{
		"~",
		"test~",
		".tmp",
		".xxx",
		"abc.tmp",
		"cvs",
		"/a/b/cvs/test.txt",
		"/a/b/.c/d",
	}
	for _, example := range falseExamples {
		if IsValidAssetPath(example) {
			t.Errorf("%q should not be a valid asset path", example)
		}
	}

	trueExamples := []string{
		"test",
		"text.txt",
		"a~b",
		"a~~~b~~c",
		"a/b/c",
		"/a/b/c",
		"a/~b/c",
		"a/b./c",
	}
	for _, example := range trueExamples {
		if !IsValidAssetPath(example) {
			t.Errorf("%q should be a valid asset path", example)
		}
	}
}

func TestIsMetaFile_TrueCase(t *testing.T) {
	examples := []string{
		"test.meta",
		"a/test.meta",
		"/a/b/test.meta",
		"test.META",
	}
	for _, example := range examples {
		if !IsMetaFile(example) {
			t.Errorf("expected true, but got false on %q", example)
		}
	}
}

func TestIsMetaFile_FalseCase(t *testing.T) {
	examples := []string{
		"test",
		"test.txt",
		"a/test",
		".meta",
		".META",
		"a/.meta",
	}
	for _, example := range examples {
		if IsMetaFile(example) {
			t.Errorf("expected false, but got true on %q", example)
		}
	}
}

func TestRemoveMetaExt(t *testing.T) {
	actual := RemoveMetaExt("test.meta")
	if actual != "test" {
		t.Errorf("expected %q, but got %q", "test", actual)
	}
}

package models

import "testing"

func Test_Result_HasContent(t *testing.T) {
	if NewResult([]string{}, []string{}).HasContent() {
		t.Error("HasContent() is true, but result is empty")
	}
	if !NewResult([]string{"a"}, []string{}).HasContent() {
		t.Error("HasContent() is false, but result has a dangling meta path")
	}
	if !NewResult([]string{}, []string{"b"}).HasContent() {
		t.Error("HasContent() is false, but result has a metaless asset path")
	}
	if !NewResult([]string{"a"}, []string{"b"}).HasContent() {
		t.Error("HasContent() is false, but result has both of a dangling meta path and a metaless asset path")
	}
}

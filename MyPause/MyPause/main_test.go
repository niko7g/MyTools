package main

import "testing"

func TestParseDelaySeconds(t *testing.T) {
	if 1 != ParseDelaySeconds("1", 10) {
		t.Errorf("出错了\n")
	}
	if 2 != ParseDelaySeconds("x2d", 10) {
		t.Errorf("出错了\n")
	}
	if 13 != ParseDelaySeconds("-1x-3d", 10) {
		t.Errorf("出错了\n")
	}
	if 5 != ParseDelaySeconds("-x-d", 5) {
		t.Errorf("出错了\n")
	}
}

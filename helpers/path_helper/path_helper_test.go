package path_helper

import (
	"testing"
)

func TestPathWithoutExt(t *testing.T) {
	expected := "path/without/ext"
	ext := ".mp4"

	value := PathWithoutExt(expected + ext)
	if value != expected {
		panic("not equal")
	}
}

func TestBasePathWithoutExt(t *testing.T) {
	long := "path/without/ext"
	expected := "ext"
	ext := ".mp4"

	value := BasePathWithoutExt(long + ext)
	if value != expected {
		panic("not equal")
	}
}

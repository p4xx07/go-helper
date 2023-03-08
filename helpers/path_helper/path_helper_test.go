package path_helper

import "testing"

func TestPathWithoutExt(t *testing.T) {
	path := "asdf/asdf/asdf.mp4"
	expected := "asdf/asdf/asdf"
	result := PathWithoutExt(path)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	path = "path/to/file.txt"
	expected = "path/to/file"
	result = PathWithoutExt(path)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	path = "path/to/folder/"
	expected = "path/to/folder/"
	result = PathWithoutExt(path)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	path = ""
	expected = ""
	result = PathWithoutExt(path)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestBasePathWithoutExt(t *testing.T) {
	path := "a/b/asdf.mp4"
	expected := "asdf"
	result := BasePathWithoutExt(path)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

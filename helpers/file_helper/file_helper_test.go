package file_helper

import (
	"os"
	"testing"
)

func TestListSorted(t *testing.T) {
	tempDir := "testdir"
	err := os.Mkdir(tempDir, 0755)
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}
	defer os.RemoveAll(tempDir)

	testFiles := []string{"file3.txt", "file1.txt", "file2.txt"}
	for _, filename := range testFiles {
		file, err := os.Create(tempDir + "/" + filename)
		if err != nil {
			t.Fatal("Failed to create test file:", err)
		}
		defer file.Close()
	}

	entries, err := ListSorted(tempDir)
	if err != nil {
		t.Fatal("Failed to list sorted entries:", err)
	}

	expectedOrder := []string{"file1.txt", "file2.txt", "file3.txt"}
	for i, entry := range entries {
		if entry.Name() != expectedOrder[i] {
			t.Errorf("Expected entry at index %d to be %s, got %s", i, expectedOrder[i], entry.Name())
		}
	}
}

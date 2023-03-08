package path_helper

import (
	"path"
	"path/filepath"
)

func PathWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func BasePathWithoutExt(fileName string) string {
	fileName = path.Base(fileName)
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

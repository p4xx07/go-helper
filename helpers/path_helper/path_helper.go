package path_helper

import (
	"path"
	"strings"
)

func PathWithoutExt(filepath string) string {
	return strings.TrimSuffix(filepath, path.Ext(filepath))
}

func BasePathWithoutExt(filepath string) string {
	return path.Base(PathWithoutExt(filepath))
}

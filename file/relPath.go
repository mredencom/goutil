package file

import (
	"path/filepath"
	"strings"
)

// 转相对路径
func RelPath(targetPath string) string {
	basePath, _ := filepath.Abs("./")
	rel, _ := filepath.Rel(basePath, targetPath)
	return strings.Replace(rel, `\`, `/`, -1)
}

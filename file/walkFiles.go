package file

import (
	"os"
	"path/filepath"
	"strings"
)

// 遍历文件，可指定后缀
func WalkFiles(targetPath string, suffixes ...string) (fileList []string) {
	if !filepath.IsAbs(targetPath) {
		targetPath, _ = filepath.Abs(targetPath)
	}
	err := filepath.Walk(targetPath, func(retPath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			fileList = append(fileList, retPath)
			return nil
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(retPath, suffix) {
				fileList = append(fileList, retPath)
			}
		}
		return nil
	})

	if err != nil {
		return
	}
	return
}

package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Dirwalk ...ディレクトリ内ファイルをを再帰的に検索し返す。
func Dirwalk(dir string, isIncludeRoot bool) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths2 := Dirwalk(filepath.Join(dir, file.Name()), isIncludeRoot)
			if !isIncludeRoot {
				for f := range paths2 {
					paths2[f] = filepath.Join(file.Name(), paths2[f])
				}
			}
			paths = append(paths, paths2...)
			continue
		}
		if isIncludeRoot {
			paths = append(paths, filepath.Join(dir, file.Name()))
		} else {
			paths = append(paths, file.Name())
		}
	}

	return paths
}

// CopyFile ...ファイルをコピーする。path1 → path2
func CopyFile(path1 string, path2 string) {
	dir := filepath.Dir(path2)
	os.MkdirAll(dir, os.ModePerm)
	raw, _ := ioutil.ReadFile(path1)
	ioutil.WriteFile(path2, raw, 0666)
}

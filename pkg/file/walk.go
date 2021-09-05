package file

import (
	"os"
	"path/filepath"
)

func (m *manager) Walk(root string, recurse bool, target string, ignore []string) ([]string, error) {
	var (
		search = []string{}
		err    error
	)

	if !recurse {
		search, err = filepath.Glob(root + target)
		if err != nil {
			return search, err
		}

		for f := range search {
			globbed := search[f][:len(search[f])-5]
			search[f] = globbed
		}

		return search, nil
	}

	err = filepath.Walk(root, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, i := range ignore {
			if fileInfo.Name() == i {
				return filepath.SkipDir
			}
		}

		if !fileInfo.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		df, err := file.Readdirnames(-1)
		if err != nil {
			return err
		}

		if matchFound(target, df) {
			search = append(search, path)
		}

		return nil
	})

	if err != nil {
		return search, err
	}

	return search, nil
}

func matchFound(match string, queryAgainst []string) bool {
	for _, t := range queryAgainst {
		if t == match {
			return true
		}
	}

	return false
}

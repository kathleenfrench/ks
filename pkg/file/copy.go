package file

import (
	"fmt"
	"io"
)

func (m *manager) CopyFile(srcPath string, targetPath string) error {
	newFile, err := m.create(targetPath)
	if err != nil {
		return fmt.Errorf("could not create file at %s - %w", targetPath, err)
	}

	defer func() {
		closeErr := newFile.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("error copying file from %s to %s - %w", srcPath, targetPath, closeErr)
		}
	}()

	input, err := m.open(srcPath)

	defer func() {
		closeErr := input.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("error copying file from %s to %s - %w", srcPath, targetPath, closeErr)
		}
	}()

	if err != nil {
		return fmt.Errorf("error copying file from %s to %s - %w", srcPath, targetPath, err)
	}

	_, err = io.Copy(newFile, input)
	if err != nil {
		return fmt.Errorf("error copying file from %s to %s - %w", srcPath, targetPath, err)
	}

	return nil
}

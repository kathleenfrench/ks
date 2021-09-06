package file

import "fmt"

func (m *manager) ReadFile(path string) ([]byte, error) {
	fileExists, err := m.FileExists(path)
	switch {
	case err != nil:
		return nil, fmt.Errorf("error reading file %s - %w", path, err)
	case !fileExists:
		return nil, fmt.Errorf("file %s does not exist", path)
	default:
		return m.read(path)
	}
}

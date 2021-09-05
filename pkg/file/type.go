package file

func (m *manager) IsFile(path string) (bool, error) {
	f, err := m.IsDirectory(path)
	if err != nil {
		return false, err
	}

	return !f, nil
}

func (m *manager) IsDirectory(path string) (bool, error) {
	d, err := m.stat(path)
	if err != nil {
		return false, err
	}

	return d.IsDir(), nil
}

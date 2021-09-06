package file

func (m *manager) FileExists(path string) (bool, error) {
	fileExists, err := m.FilepathExists(path)
	switch {
	case !fileExists:
		return false, nil
	case err != nil:
		return false, err
	default:
		return m.IsFile(path)
	}
}

func (m *manager) FilepathExists(path string) (bool, error) {
	_, err := m.stat(path)
	switch {
	case err == nil:
		return true, nil
	case m.notExistErr(err):
		return false, nil
	default:
		return false, err
	}
}

func (m *manager) DirectoryExists(path string) (bool, error) {
	directoryExists, err := m.FilepathExists(path)
	switch {
	case err != nil:
		return false, err
	case !directoryExists:
		return false, nil
	default:
		return m.IsDirectory(path)
	}
}

package file

// Remove deletes a file or directory at whatever the provided path is
func (m *manager) Remove(path string) error {
	return m.rm(path)
}

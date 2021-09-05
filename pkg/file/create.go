package file

import "fmt"

func (m *manager) Touch(path string, ps ...PermissionSetter) error {
	return m.Write(path, nil, ps...)
}

func (m *manager) Write(path string, data []byte, ps ...PermissionSetter) error {
	// give user read/write by default
	perms := setDefaults(0600)
	for _, p := range ps {
		p(&perms)
	}

	fileExists, err := m.FileExists(path)
	switch {
	case err != nil:
		return fmt.Errorf("there was an error writing to %s - %w", path, err)
	case fileExists:
		isDirectory, dirErr := m.IsDirectory(path)
		switch {
		case dirErr != nil:
			return fmt.Errorf("there was an error writing to %s - %w", path, dirErr)
		case isDirectory:
			return fmt.Errorf("%s is a directory", path)
		}
	case !fileExists:
		parentDirectory := m.dirpath(path)
		// give user read/write/execute permissions on the directory by default
		pdirErr := m.mkdir(parentDirectory, 0700)
		if pdirErr != nil {
			return fmt.Errorf("could not recursively create the directories for %s - %w", path, err)
		}
	}

	err = m.write(path, data, perms.mode)
	if err != nil {
		return fmt.Errorf("cannot write to file %s - %w", path, err)
	}

	return nil
}

func (m *manager) CreateDirectory(path string, ps ...PermissionSetter) error {
	perms := setDefaults(0700)
	for _, p := range ps {
		p(&perms)
	}

	pathExists, err := m.FilepathExists(path)
	switch {
	case pathExists:
		isFile, fileErr := m.IsFile(path)
		switch {
		case isFile:
			return fmt.Errorf("a file already exists at %s - cannot add a directory", path)
		case fileErr != nil:
			return fmt.Errorf("could not determine the type for %s - %w", path, err)
		}

		return m.chmod(path, perms.mode)
	case !pathExists:
		mkdirErr := m.mkdir(path, perms.mode)
		if mkdirErr != nil {
			return fmt.Errorf("could not create directory at %s - %w", path, err)
		}
	case err != nil:
		return fmt.Errorf("could not create directory %s - %w", path, err)
	}

	return nil
}

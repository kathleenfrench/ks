package file

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type PermissionSetter func(po *permissionSettings)

type Manager interface {
	Touch(path string, p ...PermissionSetter) error
	Write(path string, data []byte, p ...PermissionSetter) error
	FileExists(path string) (bool, error)
	FilepathExists(path string) (bool, error)
	IsFile(path string) (bool, error)
	CopyFile(srcPath string, targetPath string) error
	ReadFile(path string) ([]byte, error)
	DirectoryExists(path string) (bool, error)
	CreateDirectory(path string, p ...PermissionSetter) error
	IsDirectory(path string) (bool, error)
	Remove(path string) error
	CWD() (string, error)
	Basename(path string) string
	Homedir() (string, error)
	Walk(root string, recurse bool, target string, ignore []string) ([]string, error)
}

type manager struct {
	mkdir       func(path string, perm os.FileMode) error
	rm          func(path string) error
	write       func(name string, data []byte, mode os.FileMode) error
	dirpath     func(path string) string
	readDir     func(name string) ([]os.FileInfo, error)
	read        func(name string) ([]byte, error)
	stat        func(name string) (os.FileInfo, error)
	open        func(name string) (*os.File, error)
	create      func(name string) (*os.File, error)
	notExistErr func(err error) bool
	chmod       func(name string, mod os.FileMode) error
	cwd         func() (string, error)
	basename    func(path string) string
}

func NewManager() Manager {
	return &manager{
		mkdir:       os.MkdirAll,
		rm:          os.RemoveAll,
		write:       ioutil.WriteFile,
		dirpath:     filepath.Dir,
		readDir:     ioutil.ReadDir,
		read:        ioutil.ReadFile,
		stat:        os.Stat,
		open:        os.Open,
		create:      os.Create,
		notExistErr: os.IsNotExist,
		chmod:       os.Chmod,
		cwd:         os.Getwd,
		basename:    filepath.Base,
	}
}

type permissionSettings struct {
	mode os.FileMode
}

// SetPermissions sets the permissions mode value
func SetPermissions(p os.FileMode) PermissionSetter {
	return func(opts *permissionSettings) {
		opts.mode = p
	}
}

func setDefaults(df os.FileMode) permissionSettings {
	return permissionSettings{
		mode: df,
	}
}

func (m *manager) Homedir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return home, err
	}

	return home, nil
}

func (m *manager) CWD() (string, error) {
	return m.cwd()
}

func (m *manager) Basename(path string) string {
	return m.basename(path)
}

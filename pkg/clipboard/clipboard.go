package clipboard

import (
	"errors"
	"log"
	"os/exec"
	"runtime"
)

var (
	cpCmd     []string
	pasteCmd  []string
	noSupport bool
)

var errPlatformNotYetSupported = errors.New("ks does not support your OS yet, sry")

type Clipboard interface {
	Read() (string, error)
	Write(s string) error
}

type clip struct {
	CopyCommand  []string
	PasteCommand []string
}

func NewClipboard() Clipboard {
	return &clip{
		CopyCommand:  cpCmd,
		PasteCommand: pasteCmd,
	}
}

func (c *clip) copy() *exec.Cmd {
	switch len(cpCmd) {
	case 1:
		return exec.Command(cpCmd[0])
	default:
		return exec.Command(cpCmd[0], cpCmd[1:]...)
	}
}

func (c *clip) paste() *exec.Cmd {
	switch len(pasteCmd) {
	case 1:
		return exec.Command(pasteCmd[0])
	default:
		return exec.Command(pasteCmd[0], pasteCmd[1:]...)
	}
}

func (c *clip) Read() (string, error) {
	if noSupport {
		return "", errPlatformNotYetSupported
	}

	pasteExec := c.paste()
	out, err := pasteExec.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (c *clip) Write(s string) error {
	if noSupport {
		return errPlatformNotYetSupported
	}

	copyExec := c.copy()
	in, err := copyExec.StdinPipe()
	if err != nil {
		return err
	}

	err = copyExec.Start()
	if err != nil {
		return err
	}

	_, err = in.Write([]byte(s))
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	return copyExec.Wait()
}

func binFound(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func init() {
	switch os := runtime.GOOS; os {
	case "darwin":
		if binFound("pbcopy") {
			cpCmd = append(cpCmd, "pbcopy")
		}

		if binFound("pbpaste") {
			pasteCmd = append(pasteCmd, "pbpaste")
		}
	case "linux":
		xsel := binFound("xsel")
		xclip := binFound("xclip")

		if xsel {
			cpCmd = append(cpCmd, "xsel", "--input", "--clipboard")
			pasteCmd = append(pasteCmd, "xsel", "--output", "--clipboard")
		} else if xclip {
			cpCmd = append(cpCmd, "xclip", "-in", "-selection", "clipboard")
			pasteCmd = append(pasteCmd, "xclip", "-out", "-selection", "clipboard")
		}
	default:
		noSupport = true
		log.Fatal("OS not currently supported")
	}
}

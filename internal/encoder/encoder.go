package encoder

import (
	"fmt"
	"strings"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/pkg/clipboard"
	"github.com/kathleenfrench/ks/pkg/parse"
)

func Run(secret string, silent bool) error {
	p := parse.NewParser()
	clip := clipboard.NewClipboard()

	encoded, err := p.Encode(strings.TrimSpace(secret))
	if err != nil {
		return err
	}

	err = clip.Write(encoded)
	if err != nil {
		return err
	}

	if !silent {
		theme.Result(fmt.Sprintf("%s copied to clipboard", encoded))
	}

	return nil
}

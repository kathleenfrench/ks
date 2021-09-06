package decoder

import (
	"strings"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/pkg/clipboard"
	"github.com/kathleenfrench/ks/pkg/parse"
)

func Run(secret string, silent bool, verbose bool) error {
	p := parse.NewParser()
	clip := clipboard.NewClipboard()

	decoded, err := p.Decode(strings.TrimSpace(secret))
	if err != nil {
		return err
	}

	err = clip.Write(decoded)
	if err != nil {
		return err
	}

	if !silent {
		theme.Result(decoded)

		if verbose {
			theme.Info("> copied decoded secret to clipboard!")
		}
	}

	return nil
}

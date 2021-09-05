package ks

import (
	"os"
	"strings"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"d"},
	Example: "ks [decode, d] ZmFrZXNlY3JldA==",
	Short:   "decode a base64 encoded secret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			theme.Err("must provide a secret to decode")
			os.Exit(1)
		}

		secret := args[0]
		decoded, err := p.Decode(strings.TrimSpace(secret))
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}

		err = clip.Write(decoded)
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}

		if !silent {
			theme.Result(decoded)

			if verbose {
				theme.Info("> copied decoded secret to clipboard!")
			}
		}
	},
}

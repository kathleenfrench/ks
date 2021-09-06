package ks

import (
	"os"
	"strings"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"e"},
	Short:   "base64 encode a secret",
	Example: "ks [encode, e] fakesecret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			theme.Err("must provide a secret to encode")
			_ = cmd.Help()
			os.Exit(1)
		}

		secret := args[0]
		encoder(secret)
	},
}

func encoder(secret string) {
	encoded, err := p.Encode(strings.TrimSpace(secret))
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	err = clip.Write(encoded)
	if err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}

	if !silent {
		theme.Result(encoded)

		if verbose {
			theme.Info("> copied encoded secret to clipboard!")
		}
	}
}

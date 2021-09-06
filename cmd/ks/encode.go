package ks

import (
	"os"

	"github.com/kathleenfrench/ks/internal/encoder"
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
		err := encoder.Run(secret, silent, verbose)
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}
	},
}

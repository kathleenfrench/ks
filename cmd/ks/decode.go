package ks

import (
	"fmt"
	"os"

	"github.com/kathleenfrench/ks/internal/decoder"
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
			_ = cmd.Help()
			os.Exit(1)
		}

		if targetFile != "" {
			fmt.Println("TODO: TARGET FILE: ", targetFile)
		}

		secret := args[0]
		err := decoder.Run(secret, silent, verbose)
		if err != nil {
			theme.Err(err.Error())
			os.Exit(1)
		}
	},
}

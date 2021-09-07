package ks

import (
	"os"

	"github.com/kathleenfrench/ks/internal/decoder"
	"github.com/kathleenfrench/ks/internal/secret"
	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/kathleenfrench/ks/internal/ui"
	"github.com/kathleenfrench/ks/pkg/clipboard"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"d"},
	Example: "ks d ZmFrZXNlY3JldA==\nks d -f secret.yaml",
	Short:   "decode a base64 encoded secret or decode all secret values in a k8s yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if targetFile == "" {
			if len(args) == 0 {
				theme.Err("must provide a secret to decode")
				_ = cmd.Help()
				os.Exit(1)
			}

			secret := args[0]
			err := decoder.Run(secret, silent)
			if err != nil {
				theme.Err(err.Error())
				os.Exit(1)
			}

			return
		}

		sm := secret.NewManager()
		blob, err := sm.Parse(targetFile)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}

		ub, err := sm.DecodeData(blob)
		if err != nil {
			ui.ExitOnErr(err.Error())
		}

		theme.Info(ub.Raw)

		if !noCopy {
			clip := clipboard.NewClipboard()
			_ = clip.Write(ub.Raw)
			theme.Result("result copied to clipboard!")
		}
	},
}

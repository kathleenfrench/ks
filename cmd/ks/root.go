package ks

import (
	"os"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "a simple utility for base64 encoding secrets for k8s and copying them to the clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		switch targetFile {
		case "":
			theme.PrintLogo()
			_ = cmd.Help()
		default:
			handleFile(targetFile)
		}
	},
}

func init() {
	// flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "no std output - clipboard only mode")
	rootCmd.PersistentFlags().StringVarP(&targetFile, "file", "f", "", "target an existing secret yaml file -> ks -f <secret-filename>.yaml|.yml")

	// subcommands
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(interactiveCmd)

	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		theme.Err(err.Error())
		os.Exit(1)
	}
}

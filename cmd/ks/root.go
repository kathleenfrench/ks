package ks

import (
	"log"
	"os"

	"github.com/kathleenfrench/ks/internal/theme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "a simple utility for base64 encoding secrets for k8s and copying them to the clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		theme.PrintLogo()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

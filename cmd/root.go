package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-registry-search",
	Short: "Search images in a Docker registry",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

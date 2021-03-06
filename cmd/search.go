package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tmerello2001/docker-registry-search/search"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search images in a Docker registry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red("Search term not provided")
			os.Exit(1)
		}
		searchTerm := args[0]

		registryUrl, _ := cmd.Flags().GetString("registry")
		if len(registryUrl) < 1 {
			color.Red("Registry URL is required")
			os.Exit(1)
		}

		useHttps, _ := cmd.Flags().GetBool("https")
		imagesSrc := search.SearchImage(searchTerm, registryUrl, useHttps)
		var images []string
		for _, img := range imagesSrc {
			images = append(images, img.Refs...)
		}

		if len(images) < 1 {
			color.Yellow("Image not found")
			os.Exit(0)
		}

		prompt := promptui.Select{
			Label: fmt.Sprintf("Found %d images", len(images)),
			Items: images,
			Size:  1000,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		dockerCommand := exec.Command("docker", "pull", result)

		var stdout bytes.Buffer
		dockerCommand.Stdout = &stdout
		dockerErr := dockerCommand.Run()

		if dockerErr != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(stdout.String())

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringP("registry", "r", "", "Registry url")
	searchCmd.PersistentFlags().Bool("https", false, "Use HTTPS instead of HTTP")
}

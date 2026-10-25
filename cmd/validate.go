package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/OssiPesonen/deployment-app/internal/config"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration file",
	Long:  `Validate the application configuration file for every project defined in it.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("config.yml")
		if err != nil {
			fmt.Println("No config.yml file found. Please run 'init' before 'validate'.")
			return
		}

		config := &config.Config{}
		err = yaml.Unmarshal([]byte(data), config)
		if err != nil {
			log.Fatalf("err: %v", err)
		}

		errCount := 0

		fmt.Println("")

		// Validate config values.
		// Todo: Could be done with some form of schema that contains validation params
		if config.Version != 1 {
			fmt.Println("Configuration currently only supports version 1")
			errCount++
		}

		if len(config.Applications) == 0 {
			fmt.Println("No applications listed in config")
			errCount++
		}

		for i, app := range config.Applications {
			if app.Name == "" {
				fmt.Printf("Application %v is missing a name missing\n", i+1)
				errCount++
			}

			if app.Artifact.Repository == "" {
				fmt.Printf("Application %v is missing repository info\n", i+1)
				errCount++
			} else {
				repo := strings.Split(app.Artifact.Repository, "/")
				if len(repo) != 2 {
					fmt.Printf("Application repository in incorrect format. Please provide it as 'username/repository' string\n")
					errCount++
				}
			}

			if app.Artifact.Selector == "" {
				fmt.Printf("Application %v is missing artifact selector\n", i+1)
				errCount++
			}

			if app.Artifact.Type != "github-release" {
				fmt.Printf("Application %v artifact type incorrect. Only 'github-release' currently supported\n", i+1)
				errCount++
			}

			if app.Artifact.TokenEnv == "" {
				fmt.Printf("Application %v1 artifact no token provided. Unable to enable fetching package\n", i+1)
				errCount++
			}
		}

		fmt.Println("")

		if errCount > 0 {
			fmt.Printf("✗ Validation failed. Errors: %v\n", errCount)
		} else {
			fmt.Println("✓ Configuration is valid.")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

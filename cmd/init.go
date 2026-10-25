/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/OssiPesonen/deployment-app/internal/config"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application and first configuration",
	Long: ` 
	Used to initialize the application for the first time with an basic
	configuration for Github
	.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("config.yml")
		if err == nil {
			fmt.Println("Initialization already done. If you wish to re-do it, remove the config.yml file and run the command again")
			return
		}

		var project string
		var repository string

		fmt.Println("Give application/project name:")
		fmt.Scan(&project)
		fmt.Println("Give GH repository path (username/repository):")
		fmt.Scan(&repository)

		applications := []config.Application{
			{
				Name: project,
				Artifact: config.ApplicationArtifact{
					Repository: "OssiPesonen/go-builder-app",
					TokenEnv:   "",
					Selector:   "selector",
					Type:       "github-release",
				},
			},
		}

		config := &config.Config{
			Version:      1,
			Applications: applications,
		}

		y, yErr := yaml.Marshal(config)
		if yErr != nil {
			log.Fatalf("error: %v", err)
		}

		err = os.WriteFile("config.yml", y, 0o644)
		if err == nil {
			fmt.Println("Initialization completed")
		} else {
			fmt.Println("An error occurred")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

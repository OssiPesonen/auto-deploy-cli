package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deployment-app",
	Short: "Lightweight deployment application",
	Long: `This is a lightweight deployment app that
	can automate your application installments on your
	on-premise servers without the need for cloud platforms
	or complex installation setups.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deployment-app.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

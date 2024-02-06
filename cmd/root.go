package cmd

import (
	"init/internal/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your application",
}

func Execute() {
	utils.SetupConfig()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

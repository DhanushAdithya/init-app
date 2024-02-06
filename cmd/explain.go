package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("explain called")
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	explainCmd.Flags().BoolP("last", "l", true, "Explain the last command")
	explainCmd.Flags().BoolP("session", "s", false, "Explain the session")

	viper.BindPFlag("last", explainCmd.Flags().Lookup("last"))
	viper.BindPFlag("session", explainCmd.Flags().Lookup("session"))

	viper.SetDefault("last", true)
	viper.SetDefault("session", false)
}

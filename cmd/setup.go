package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("shell", viper.GetString("shell"))
		viper.Set("history", viper.GetString("history"))
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringP("shell", "s", "bash", "The shell to use")
	setupCmd.Flags().StringP("history", "H", "", "Path of the history file")

	viper.BindPFlag("shell", setupCmd.Flags().Lookup("shell"))
	viper.BindPFlag("history", setupCmd.Flags().Lookup("history"))

	viper.SetDefault("shell", "bash")
	viper.SetDefault("history", "")
}

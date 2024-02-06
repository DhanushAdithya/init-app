package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Executes the user prompt with necessary commands",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		fmt.Println(prompt)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	doCmd.Flags().BoolP("verbose", "v", false, "Prints the output of the command to the console")
	doCmd.Flags().BoolP("interactive", "i", false, "Runs the command in interactive mode")
	doCmd.Flags().BoolP("solo", "s", false, "Runs the command in isolated mode")
	doCmd.Flags().BoolP("memloss", "m", false, "Runs the command without the previous history")

	viper.BindPFlag("verbose", doCmd.Flags().Lookup("verbose"))
	viper.BindPFlag("interactive", doCmd.Flags().Lookup("interactive"))
	viper.BindPFlag("solo", doCmd.Flags().Lookup("solo"))
	viper.BindPFlag("memloss", doCmd.Flags().Lookup("memloss"))

	viper.SetDefault("verbose", false)
	viper.SetDefault("interactive", false)
	viper.SetDefault("solo", false)
	viper.SetDefault("memloss", false)
}

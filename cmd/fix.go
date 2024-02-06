package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("history") == "" {
			fmt.Println("History file does not exist")
			return
		}
		history, err := os.ReadFile(viper.GetString("history"))
		if err != nil {
			fmt.Println(err)
			return
		}
		historyCmds := strings.Split(string(history), "\n")
		lastCmd := historyCmds[len(historyCmds)-3]
		fmt.Println(lastCmd)
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}

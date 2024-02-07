package cmd

import (
	"fmt"
	"init/internal/fetch"
	"init/internal/utils"
	"os"
	"os/exec"
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
		lcmd := exec.Command(lastCmd)
		err = lcmd.Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				res, err := fetch.Fetch(fmt.Sprintf(utils.FixPrompt, lastCmd, exitError.ExitCode()))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(res.Response)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}

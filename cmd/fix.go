package cmd

import (
	"fmt"
	"init/internal/fetch"
	"init/internal/tui"
	"init/internal/utils"
	"os"
	"sync"

	// "os/exec"
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
		historyCmds := strings.Split(strings.TrimSpace(string(history)), "\n")
		lastCmd := historyCmds[len(historyCmds)-2]
		// lcmd := exec.Command(lastCmd)
		// err = lcmd.Run()
		// if err != nil {
		// 	if exitError, ok := err.(*exec.ExitError); ok {
		res, err := fetch.Fetch(fmt.Sprintf(utils.FixPrompt, lastCmd, 1))
		var wg sync.WaitGroup
		wg.Add(1)
		response := make(chan struct{})
		go func() {
			defer wg.Done()
			tui.RenderLoad("Generating solution to fix the "+"`"+lastCmd+"` command", response)
		}()
		if err != nil {
			close(response)
			wg.Wait()
			fmt.Println(err)
			return
		}
		close(response)
		wg.Wait()
		fmt.Println(res.Response)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}

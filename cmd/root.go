package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"leinao-os-adapter/pkg/utils"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "leinao-os-adapter",
	Short: "adapted leinao-os interface",
	Long: `  ██         ██                                                                         ██                     ██                 
 ░██        ░░                                                                         ░██           ██████   ░██                 
 ░██  █████  ██ ███████   ██████    ██████         ██████   ██████        ██████       ░██  ██████  ░██░░░██ ██████  █████  ██████
 ░██ ██░░░██░██░░██░░░██ ░░░░░░██  ██░░░░██ █████ ██░░░░██ ██░░░░  █████ ░░░░░░██   ██████ ░░░░░░██ ░██  ░██░░░██░  ██░░░██░░██░░█
 ░██░███████░██ ░██  ░██  ███████ ░██   ░██░░░░░ ░██   ░██░░█████ ░░░░░   ███████  ██░░░██  ███████ ░██████   ░██  ░███████ ░██ ░ 
 ░██░██░░░░ ░██ ░██  ░██ ██░░░░██ ░██   ░██      ░██   ░██ ░░░░░██       ██░░░░██ ░██  ░██ ██░░░░██ ░██░░░    ░██  ░██░░░░  ░██   
 ███░░██████░██ ███  ░██░░████████░░██████       ░░██████  ██████       ░░████████░░██████░░████████░██       ░░██ ░░██████░███   
░░░  ░░░░░░ ░░ ░░░   ░░  ░░░░░░░░  ░░░░░░         ░░░░░░  ░░░░░░         ░░░░░░░░  ░░░░░░  ░░░░░░░░ ░░         ░░   ░░░░░░ ░░░    
`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(utils.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func tip() {
	usageStr := `欢迎使用 ` + utils.Green(`leinao-os-adapter`) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

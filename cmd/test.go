package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Get Variables Enviroment",
	Long:  `Execute Command: ./todo test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("My HOME:", os.Getenv("HOME"))
		fmt.Println("My GOPATH:", os.Getenv("GOPATH"))
		fmt.Println("My PATH:", os.Getenv("PATH"))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

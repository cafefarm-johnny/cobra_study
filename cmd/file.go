package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:     "file",
	Short:   "File handling command",
	Long:    "Which reading, writing for files",
	Example: "file --help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please enter command.")
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	//fileCmd.Flags().StringVarP(&filePath, "read", "r", "", "read for file")
	//fileCmd.Flags().StringVarP(nil, "list", "l", "", "list up files")
}

func fileExist(path string) bool {
	_, file := filepath.Split(path)
	if len(file) < 1 {
		return false
	}
	if filepath.Ext(file) == "" {
		return false
	}
	return true
}

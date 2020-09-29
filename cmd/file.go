package cmd

import (
	"cobra_sample/utils/delimiter"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var filePath string
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
}

func isDir(path string) bool {
	_, file := filepath.Split(path)
	if len(file) < 1 {
		return false
	}
	if filepath.Ext(file) != delimiter.Blank {
		return false
	}
	return true
}

func fileExist(path string) bool {
	_, file := filepath.Split(path)
	if len(file) < 1 {
		return false
	}
	if filepath.Ext(file) == delimiter.Blank {
		return false
	}
	return true
}

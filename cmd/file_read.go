package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var filePath string
var readCmd = &cobra.Command{
	Use:     "read",
	Short:   "File reading command",
	Long:    "Which reading for files",
	Example: "file read -p ~/downloads/test.txt",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !fileExist(filePath) {
			return fmt.Errorf("\"%s\" is directory or there is no file", filePath)
		}
		err := readFile()
		if err != nil {
			return errors.New("cannot open file")
		}

		return nil
	},
}

func init() {
	fileCmd.AddCommand(readCmd)

	readCmd.PersistentFlags().StringVarP(&filePath, "url", "u", "", "file path for reading")
}

func readFile() error {
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}

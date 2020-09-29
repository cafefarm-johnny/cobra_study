package cmd

import (
	"cobra_sample/utils/delimiter"
	"cobra_sample/utils/nums"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "File list up command",
	Long:    "Which list up for files",
	Example: "file list -p ~/downloads",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !isDir(filePath) {
			return fmt.Errorf("\"%s\" is not directory", filePath)
		}
		err := list(filePath)
		if err != nil {
			return errors.New("cannot list up files")
		}
		return nil
	},
}

func init() {
	fileCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&filePath, "url", "u", delimiter.Blank, "file path for list up")
}

func list(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range files {
		printFormat := "%-100v %-30v %v " + delimiter.LineAlignment
		size := strconv.FormatInt(file.Size(), nums.Decimal) + delimiter.Byte
		time := strings.Replace(file.ModTime().Format(time.RFC3339), delimiter.Time, delimiter.WhiteSpace, 1)

		fmt.Printf(printFormat, file.Name(), size, time)
	}

	return nil
}

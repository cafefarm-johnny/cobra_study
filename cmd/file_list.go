package cmd

import (
	"cobra_sample/utils/nums"
	"cobra_sample/utils/separate"
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
	Example: "file list --path [path]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !isDir(filePath) {
			return fmt.Errorf("\"%s\" is not directory", filePath)
		}
		err := listUp(filePath)
		if err != nil {
			return errors.New("cannot list up files")
		}
		return nil
	},
}

func init() {
	fileCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&filePath, "path", "p", separate.Blank, "file path for list up")
}

func listUp(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range files {
		printFormat := "%-100v %-30v %v " + separate.LineAlignment
		size := strconv.FormatInt(file.Size(), nums.Decimal) + separate.Byte
		time := strings.Replace(file.ModTime().Format(time.RFC3339), separate.Time, separate.WhiteSpace, 1)

		fmt.Printf(printFormat, file.Name(), size, time)
	}

	return nil
}

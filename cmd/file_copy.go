package cmd

import (
	"cobra_sample/utils/separate"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
)

var copyPaths = make([]string, 2, 2)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Short:   "File copy command",
	Long:    "Which copy for files",
	Example: "file copy --from [from with filename] --to [to with filename]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(copyPaths) < 2 || len(copyPaths[0]) < 1 || len(copyPaths[1]) < 1 {
			return errors.New("please input path from, to")
		}

		from := copyPaths[0]
		if !fileExist(from) {
			return fmt.Errorf("\"%s\" is directory or there is no file", from)
		}

		to := copyPaths[1]
		err := copyFile(from, to)
		if err != nil {
			return errors.New("cannot copy file")
		}

		endpoint := filepath.Dir(to)
		fmt.Println(endpoint)

		err = listUp(endpoint)
		if err != nil {
			return errors.New("cannot list up files")
		}

		return nil
	},
}

func init() {
	fileCmd.AddCommand(copyCmd)

	copyCmd.PersistentFlags().StringVarP(&copyPaths[0], "from", "f", separate.Blank, "file path of from")
	copyCmd.PersistentFlags().StringVarP(&copyPaths[1], "to", "t", separate.Blank, "file path of to")
}

func copyFile(from string, to string) error {
	bytes, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(to, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

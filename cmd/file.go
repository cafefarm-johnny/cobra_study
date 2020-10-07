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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("fileCmd > PersistentPreRun > args : ", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		subCmdAction()
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}

func subCmdAction() {
	fmt.Println("What are you want?")
	fmt.Println("1. list up")
	fmt.Println("2. copy")
	fmt.Println("3. read")
	fmt.Print("> ")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		listUpAction()
		break
	case 2:
		copyAction()
		break
	case 3:
		readAction()
		break
	default:
		fmt.Println("what?")
		subCmdAction()
		break
	}
}

func listUpAction() {
	fmt.Println("Please enter directory path")
	fmt.Println("ex) /Users/johnnyuhm/Downloads")
	fmt.Print("> ")

	var flag string
	fmt.Scanln(&flag)

	// file list --path ... 커맨드로 rootCmd(메인 함수)를 실행
	rootCmd.SetArgs([]string{"file", "list", fmt.Sprintf("--path=%s", flag)})
	rootCmd.Execute()
}

func copyAction() {
	fmt.Println("Please enter copy path \"[from] [to]\"")
	fmt.Println("ex) [from] /Users/johnnyuhm/Downloads/canvas.html")
	fmt.Println("ex) [to] /Users/johnnyuhm/Downloads/(copied)canvas.html")
	fmt.Print("> ")

	var from string
	var to string
	fmt.Scan(&from, &to)

	rootCmd.SetArgs([]string{"file", "copy", fmt.Sprintf("--from=%s", from), fmt.Sprintf("--to=%s", to)})
	rootCmd.Execute()
}

func readAction() {
	fmt.Println("Please enter read file path")
	fmt.Println("ex) /Users/johnnyuhm/Downloads/canvas.html")
	fmt.Println("> ")

	var path string
	fmt.Scan(&path)

	rootCmd.SetArgs([]string{"file", "read", fmt.Sprintf("--path=%s", path)})
	rootCmd.Execute()
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

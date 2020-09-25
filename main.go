package main

import "cobra_sample/cmd"

// 실행방법
// 1. go run $project_folder/main.go [options]
// 2. go run $project_folder/main.go Johnny [options]
func main() {
	cmd.Execute()
}

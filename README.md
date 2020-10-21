# COBRA_SAMPLE

Go Cobra CLI 라이브러리 예제를 보며 만들어 본 간단한 파일 커맨드 CLI 프로그램입니다.

링크: https://github.com/spf13/cobra

## How to use

### Build

1. go install cobra_sample

- output: ${GOPATH}/bin/cobra_sample

### Execute

1. cobra_sample [command] [subcommand] [flags] args...

2. go run ${project}/main.go [command] [subcommand] [flags] args... 

### Help

1. cobra_sample --help

### Config

1. cobra_sample --config [path of CLI config file]

### Version

1. (flag) cobra_sample --version

2. (command) cobra_sample version

### File Command

1. cobra_sample file --help

#### List Up

1. cobra_sample file list --path [path]

#### Read

1. cobra_sample file read --path [path with filename]

#### Copy

1. cobra_sample file copy --from [path with filename] --to [path with filename]

### Progress Bar

1. cobra_sample progress --help
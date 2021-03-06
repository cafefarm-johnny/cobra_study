package cmd

import (
	"cobra_sample/utils/separate"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	host string
	port int
)

const version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "cobra_sample",
	Short:   "This is short comment.",
	Long:    `This is long comment.`,
	Version: version, // 버전 플래그(--version)
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfgFile)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main.go 실행 시 cmd 패키지 안에 init()을 실행하여 커맨드를 파싱한다.
// 커맨드 파싱 완료 후 Run에 정의된 함수를 호출한다.
func init() {
	cobra.OnInitialize(initConfig)

	// 외부 설정 파일 로드 플래그
	// ex) --config ./myConfig.cli.yml
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "(test)default config", "config file (default is $HOME/.cli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// viper 라이브러리를 이용해 외부 설정 파일 로딩처리
func initConfig() {
	if cfgFile != separate.Blank {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using Config File: ", viper.ConfigFileUsed())
	}

	// config 파일 내 요소 접근하기
	host = viper.GetString("server.host")
	port = viper.GetInt("server.port")
	fmt.Println("> host : ", host)
	fmt.Println("> port : ", port)
}

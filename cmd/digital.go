package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var digitalCmd = &cobra.Command{
	Use:     "digital",
	Short:   "전자제품 주문 시스템임",
	Long:    "전자제품 주문 시스템입니다.",
	Example: "digital --config test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("전자제품 주문 시스템을 이용해주셔서 감사합니다.")
	},
}

func init() {
	rootCmd.AddCommand(digitalCmd)
}

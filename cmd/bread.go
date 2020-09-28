package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var breadCmd = &cobra.Command{
	Use:     "bread",
	Short:   "빵 주문 시스템임",
	Long:    "빵 주문 시스템입니다.",
	Example: "bread --config test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("빵 주문 시스템을 이용해주셔서 감사합니다.")
		fmt.Printf("다음 제품을 구매하시는게 맞으신가요? \"%s\"\n", product)
	},
}

var product string

func init() {
	rootCmd.AddCommand(breadCmd)
	breadCmd.PersistentFlags().StringVar(&product, "product", "cream bread", "A help for product.")
}

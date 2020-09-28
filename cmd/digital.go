package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var digitalCmd = &cobra.Command{
	Use:     "digital",
	Short:   "전자제품 주문 시스템임",
	Long:    "전자제품 주문 시스템입니다.",
	Example: "digital --config test",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("제품을 입력해주세요")
		}
		if !validation(args[0]) {
			return fmt.Errorf("\"%s\" 제품은 판매 불가 상품입니다", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("전자제품 주문 시스템을 이용해주셔서 감사합니다.")
	},
}

var notForSaleItem = "냉장고"

func init() {
	rootCmd.AddCommand(digitalCmd)
}

func validation(product string) bool {
	if strings.Contains(product, notForSaleItem) {
		return false
	}
	return true
}

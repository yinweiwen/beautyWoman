package ai

import (
	"fmt"
	"testing"
)

func TestBaiduAi_Recognize(t *testing.T) {
	ai := NewBaiduAi()
	ret, err := ai.Recognize("../../download/fix/pic2.jpg")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range ret.Result {
			fmt.Println(item.Keyword)
		}
	}
}

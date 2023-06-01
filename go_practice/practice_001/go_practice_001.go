/*
	課題1
		1(int), "2"(string), 10, "11"のスライスをループして、一桁なら01のように0をいれて2桁にして表示する。
		Go better play groundで動作確認するといい。スライスをループ→fmt.Println()で表示というイメージ。
*/

package practice_001

import (
	"fmt"
	"strconv"
)

func Func_001() {
	sl := []interface{}{1, "2", 10, "11"}

	for _, v := range sl {
		switch v := v.(type) {
		case int:
			fmt.Printf("%02d\n", v) // 0埋め
		case string:
			// 数値に変換
			if num, err := strconv.Atoi(v); err == nil {
				fmt.Printf("%02d\n", num)
			}
		}
	}
}

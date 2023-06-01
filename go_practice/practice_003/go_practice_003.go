/*
	課題3
		type MyIntSlice []intでIntスライスと互換のある独自型を作り、重複を排除するメソッドを実装する。
*/

package practice_003

import "fmt"

type MyIntSlice []int

func (m MyIntSlice) Unique() MyIntSlice {
	// 重複チェック関数
	keys := make(map[int]bool) // 重複チェックmap
	res := MyIntSlice{}        // 返却データ初期化
	for _, v := range m {
		if !keys[v] {
			// keysがfalseの場合
			res = append(res, v)
			keys[v] = true // 重複を避けるため、trueとする
		}
	}
	return res
}

func Func_003() {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]
}

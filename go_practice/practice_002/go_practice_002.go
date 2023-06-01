/*
	課題2
		mapのvalueからkeyを取得する問題。引数にmapとvalueを受け取り、対応するkeyがあればkeyを返却する関数を実装する。
		mapの型はkeyをint, valueをstringにすること。もし対応するkeyがなければerrorを返すこと。
*/

package practice_002

import (
	"errors"
	"fmt"
)

func getKey(m map[int]string, target string) (int, error) {
	// mapからkeyを取得し、呼び出し元に返却する関数。keyがない場合はerrorを返却する。
	for k, v := range m {
		if v == target {
			return k, nil
		}
	}
	return 0, errors.New("Keyが見つかりません")
}

func result(k int, err error) {
	// 引数のerrの値を判定して、標準出力する関数
	if err == nil {
		fmt.Println(k)
	} else {
		fmt.Println(err)
	}
}

func Func_002() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	k, err := getKey(m, "A")
	result(k, err)
	k, err = getKey(m, "B")
	result(k, err)
	k, err = getKey(m, "C")
	result(k, err)
	k, err = getKey(m, "D")
	result(k, err)
}

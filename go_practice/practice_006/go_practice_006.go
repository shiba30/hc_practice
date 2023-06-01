/*
	課題6
*/

package practice_006

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return "hoge"
}

func Func_006() {
	var m MyInt = 3
	fmt.Println(m) // hogeと出力させるように修正せよ。ただしmain関数に変更を加えないこと。
}

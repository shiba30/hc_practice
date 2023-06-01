/*
	課題4
		このコードの実行結果はどうなるか予想してください。
		実際に実行してなぜこのような挙動になるのかを説明してください。
		また、正しいと思われる挙動にするための修正をしてください。
*/

package practice_004

import "fmt"

type User struct {
	Name string
	Age  int
}

func Func_004() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	/*
		for _, user := range users {
			user.Age = 44
		}
	*/

	for i := range users {
		users[i].Age = 44
	}

	fmt.Printf("%v\n", users)
	/*
		実行結果はどうなる？
			出力結果 [{tarou 33} {zirou 22} {itirou 11}] （26行目の処理が更新が反映されていない）
	*/
	/*
		なぜこのような挙動になるのか？
			修正前の処理では、「_, user := range users」で、userはusersのスライスの要素をコピーしてループしている。
			要素の値参照ではなく新たな変数（コピー）をループさせているため、値を代入しても元のusersスライスには影響が出ないため、出力結果が元のスライスの値となる。
			修正した処理では、for i := range usersでループをさせ、iはusersのインデックスで、i番目の要素のAgeを更新している。
			users[i] は users の直接的な参照であるため、値更新が成功する。
	*/
}

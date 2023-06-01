/*
	課題5
		以下のコードに記号を2つだけ追加して、コメントの挙動になるように修正すること
*/

package practice_005

import "fmt"

var appConfig = Config{Env: "test"}

type Config struct {
	Env string
}

func getConfig() *Config {
	return &appConfig
}

func Func_005() {
	c := getConfig()
	c.Env = "production"

	fmt.Println(c.Env)         // production
	fmt.Println(appConfig.Env) // testではなくproducionへ
}

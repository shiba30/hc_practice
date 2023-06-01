/*
	課題7
		以下のようなjsonが複数あるログがある。(sample.log)ログは1行に1つのJSONがある。このログをGoの構造体に変換してDB(PostgreSQL)にInsertする。

		{
			"user": {
				"age": 22,
				"name": "tarou",
				"role": "student"
		},
			"dist": "PostgreSQL",
			"level": "info",
			"msg": "test",
			"src": "backend",
			"time": "2021-08-01T00:05:05Z"
		}

		PostgreSQLはDockerコンテナで作ること
		forでループしてInsertすること
		トランザクションを実装すること（全部Insert成功 or 全部失敗を補償する）1つでもエラーが起きたらrollbackすればよい。
		usersテーブルのカラムは id(pk),age(integer), name(varchar(500)), role(char(15))にすること
		dist, levelなどのデータは使用しない
		go run main.go sample.logのようにファイルを引数で渡せること
		複数の引数を渡すとエラー終了すること
		dbの中身の確認はtable plusを使うと良い
		↓をコピーして sample.logを作成する

		{"user":{"age":22,"name":"tarou","role":"student"},"dist":"PostgreSQL","level":"info","msg":"test","src":"backend","time":"2021-08-01T00:05:05Z"}
		{"user":{"age":23,"name":"zirou","role":"student"},"dist":"PostgreSQL","level":"info","msg":"test","src":"backend","time":"2021-08-01T00:05:05Z"}
		{"user":{"age":24,"name":"saburou","role":"student"},"dist":"PostgreSQL","level":"info","msg":"test","src":"backend","time":"2003-08-01T00:05:05Z"}
		{"user":{"age":25,"name":"mike","role":"mentor"},"dist":"PostgreSQL","level":"warn","msg":"test","src":"backend","time":"2011-08-01T00:05:05Z"}
		{"user":{"age":26,"name":"make","role":"mentor"},"dist":"PostgreSQL","level":"warn","msg":"test","src":"backend","time":"2020-08-01T00:05:05Z"}
*/

package practice_007

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type LogData struct {
	User User `json:"user"`
	// dist, levelなどのデータは使用しないので、定義しない
}

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "test"
	DB_PASS  = "password"
	DB_NAME  = "practice"
	DB_TABLE = "users"
)

func Func_007() {
	var fname string

	// 引数取得
	flag.StringVar(&fname, "f", "", "File read")
	flag.Parse()
	if fname == "" {
		fmt.Println("Error: no argument specified.")
		os.Exit(1)
	}
	if flag.NArg() != 0 {
		fmt.Println("Error: multiple arguments.")
		os.Exit(1)
	}

	filePath := fmt.Sprintf("%s/%s", "practice_007", fname)

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 一行ごとにjsonデータをlogDataへ変換
	var logDataList []LogData
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var logData LogData
		err := json.Unmarshal([]byte(scanner.Text()), &logData)
		if err != nil {
			panic(err)
		}
		logDataList = append(logDataList, logData)
	}

	// connect postgres
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// insert
	query := fmt.Sprintf("INSERT INTO %s (age, name, role) VALUES ($1, $2, $3)", DB_TABLE)
	for _, logData := range logDataList {
		_, err := tx.Exec(query, logData.User.Age, logData.User.Name, logData.User.Role)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	// commit
	err = tx.Commit()
	if err != nil {
		fmt.Println("Could not commit transaction:", err)
		tx.Rollback()
		os.Exit(1)
	}
}

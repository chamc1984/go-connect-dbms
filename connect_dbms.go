package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	flags "github.com/jessevdk/go-flags"

	_ "github.com/lib/pq"
)

type TABLE1 struct {
	ID   uint
	NAME string
}

var opts struct {
	// Example of a required flag
	Dbms string `short:"d" long:"database" description:"connect target dbms.[postgres/mysql]" required:"true"`
}

func main() {

	// 引数取得
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println(err)
	}

	// SQL読み込み
	f, err := os.Open("select.sql")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	// DB接続
	var arg1, arg2 string
	switch opts.Dbms {
	case "postgres":
		arg1 = "postgres"
		arg2 = "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
	case "mysql":
		arg1 = "mysql"
		arg2 = "root:password@tcp(127.0.0.1:3306)/database1"
	default:
		fmt.Println("invalid dbms.")
		return
	}
	db, err := sql.Open(arg1, arg2)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	// SELECT
	rows, err := db.Query(string(b))
	if err != nil {
		fmt.Println(err)
	}

	// 結果出力
	var es []TABLE1
	for rows.Next() {
		var e TABLE1
		rows.Scan(&e.ID, &e.NAME)
		es = append(es, e)
	}
	fmt.Printf("%v", es)
}

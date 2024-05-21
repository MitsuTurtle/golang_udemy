package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	/*
		// C
		cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
		_, err := Db.Exec(cmd, "Nancy", 20)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		// R
		// 特定のデータを取得
		cmd := "SELECT * FROM persons WHERE age = $1"
		// QueryRow 1レコード取得
		row := Db.QueryRow(cmd, 20)
		var p Person
		err := row.Scan(&p.Name, &p.Age)
		if err != nil {
			// データがなかったら
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Panicln(err)
				// それ以外のエラー
			}
		}
		fmt.Println(p.Name, p.Age)

		// 複数のデータを取得
		cmd = "SELECT * FROM persons"
		// Queryは条件に合うものと全て取得
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		// structを作成
		var pp []Person
		for rows.Next() {
			var p Person
			// scan データ追加
			err := rows.Scan(&p.Name, &p.Age)
			// 一つずつエラーハンドリングver
			if err != nil {
				log.Panicln(err)
			}
			pp = append(pp, p)
		}
		// まとめてエラーハンドリングver
		err = rows.Err()
		if err != nil {
			log.Fatalln(err)
		}
		// 表示
		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	/*
		// U
		// データの更新
		cmd := "UPDATE persons SET age = $1 WHERE name = $2"
		_, err := Db.Exec(cmd, 25, "Nancy")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// D
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConnect() {
	var err error
	// DBへの接続を行う
	db, err = sql.Open("mysql", "test_user:password@tcp(127.0.0.1:8000)/testDB?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("データベース接続成功")
}

// バイト列をUTF-8文字列に変換する関数
func bytesToString(b []byte) string {
	return string(b)
}

// usersテーブルのデータを取得する関数
func getRows() ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, colName := range columns {
			value := columnValues[i] // valueを設定
			// バイト列の場合、UTF-8に変換
			if bytesValue, ok := value.([]byte); ok {
				result[colName] = bytesToString(bytesValue) // 変換処理
			}
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	dbConnect()
	r := gin.Default()

	// ルートURL ("/") に対するGETリクエストをハンドル
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// usersテーブルのデータを取得するGETエンドポイント
	r.GET("/getUsers", func(c *gin.Context) {
		results, err := getRows()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		log.Printf("取得したデータ: %+v", results)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, gin.H{
			"data": results,
		})
	})

	// 8000ポートでサーバーを起動
	r.Run(":8010")
}

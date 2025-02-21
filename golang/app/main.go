package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/gin-gonic/gin"
	"database/sql"
	//"gorm.io/gorm"
	//"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	// DBへの接続を行う
	db, err := sql.Open("mysql", "test_user:password@tcp(127.0.0.1:3306)/testDB")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})`
	// エラーが発生した場合、エラー内容を表示
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// 接続に成功した場合、「db connected!!」と表示する
	fmt.Println("db connected!!")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("データベース接続成功")


	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートURL ("/") に対するGETリクエストをハンドル
	r.GET("/", func(c *gin.Context) {
		// JSONレスポンスを返す
		c.JSON(200, gin.H{
		"message": "Hello World",
		})
	})

	// 8080ポートでサーバーを起動
	r.Run(":8000")
}


// DBを起動させる
// func dbInit() *gorm.DB {
// 	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
// 	dsn := fmt.Sprintf(`%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True`, 
//             os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
// 	// DBへの接続を行う
// 	db, err := sql.Open("mysql",dsn)
// 	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// エラーが発生した場合、エラー内容を表示
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// 接続に成功した場合、「db connected!!」と表示する
// 	fmt.Println("db connected!!")
// 	err = db.Ping()
//     if err != nil {
//         log.Fatal(err)
//     }

// 	return db
// }

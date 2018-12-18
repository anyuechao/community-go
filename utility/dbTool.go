package utility

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//func init() {
//	d,err := sql.Open("mysql","root:zhq1346221@tcp(localhost:3306)/merlot_test?charset=utf8")
//	db = d
//	if err != nil {
//		fmt.Println(err.Error())
//	}else {
//		fmt.Println("数据库连接成功")
//	}
//}
//func InsertData(tableName string,)  {
//	fmt.Println("=====",db)
//	stmt,err:= db.Prepare("INSERT INTO user(name) VALUES (?)")
//	//stmt,err := db.Prepare("INSERT INTO user(name) VALUES (?)")
//	defer stmt.Close()
//	stmt.Exec("zhq")
//	fmt.Println(err)
//}
//func FetchData()  {
//	rows,err := db.Query("SELECT name  FROM user")
//	if err != nil {
//		fmt.Println("查表错误:",err.Error())
//	}else {
//		for rows.Next() {
//			fmt.Println("===")
//			var name string
//			rows.Scan(&name)
//			fmt.Println(name)
//		}
//	}
//}

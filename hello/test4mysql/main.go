package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (err error) {
	var dsn = "root:123456@tcp(182.92.170.252:3308)/news"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql打开失败")
		return err
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("数据库ping失败")
		return err
	}
	return nil

}
func main() {
	InitDB()
	InsertToMysql("Jony", "this is Jony", 70, "Jony")
	RunServer()
}

type TencentNews struct {
	Title, Content, Author string
	Pubtime                int
}

func InsertToMysql(author string, content string, pubtime int, title string) {
	// DSN:Data Source Name

	stmt, err := DB.Prepare("insert into items values (null,?,?,?,? )")
	CheckErr(err, "字符串预处理失败")

	res, err := stmt.Exec(author, content, pubtime, title)
	CheckErr(err, "插入执行失败")

	id, err := res.LastInsertId()
	CheckErr(err, "id失败")
	fmt.Println(id)

	//defer db.Close() // 注意这行代码要写在上面err判断的下面
}
func QueryFromMysql(id int) TencentNews {

	queryStr := "select * from items where id =?"
	var iauthor string
	var icontent string
	var ipubtime int
	var ititle string
	err := DB.QueryRow(queryStr, id).Scan(&id, &iauthor, &icontent, &ipubtime, &ititle)

	if err != nil {
		fmt.Println("查询失败")
		panic(err)
	}
	return TencentNews{Title: ititle, Author: iauthor, Content: icontent, Pubtime: ipubtime}

}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	q := r.URL.Query()

	if len(q) != 0 {

		id, err := strconv.Atoi(q.Get("id"))
		if err != nil {
			panic(err)
		}

		fmt.Println(QueryFromMysql(id))
		fmt.Fprintln(w, QueryFromMysql(id))

	}
	//content =
}

func RunServer() {

	fmt.Println("Server already open")
	http.HandleFunc("/get", GetHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Error")
		return
	}
}
func CheckErr(e error, statement string) {
	fmt.Println(statement)
	panic(e)
}

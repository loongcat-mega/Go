package Mysql

import (
	"database/sql"
	"fmt"

	. "git.woa.com/qingruixu/News"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (err error) {
	var dsn = "root:123456@tcp(182.92.170.252:3308)/news"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("sql打开失败 %v", err)
	}
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("数据库ping失败 %v", err)
	}
	return nil

}

func InsertToMysql(author string, content string, pubtime int, title string) error {
	// DSN:Data Source Name

	stmt, err := DB.Prepare("insert into items (author,content,pubtime,title) VALUES (?,?,?,? )")
	if err != nil {
		return fmt.Errorf("字符串预处理失败 %v", err)

	}
	defer stmt.Close()
	res, err := stmt.Exec(author, content, pubtime, title)
	if err != nil {
		return fmt.Errorf("插入执行失败 %v", err)
	}
	_ = res
	return nil
	//defer db.Close() // 注意这行代码要写在上面err判断的下面
}

func QueryFromMysql(id int) (TencentNews, error) {

	queryStr := "select * from items where id =?"

	//var tencentnews TencentNews
	var iauthor string
	var icontent string
	var ipubtime int
	var ititle string
	var iid int

	err := DB.QueryRow(queryStr, id).Scan(&iid, &iauthor, &icontent, &ipubtime, &ititle)
	if err != nil {
		fmt.Println("单行查询失败")
		if err == sql.ErrNoRows {
			return TencentNews{}, fmt.Errorf("No row found of id : %d", id)
		}
		return TencentNews{}, fmt.Errorf("query failed %v ", err)

	}
	return TencentNews{Author: iauthor, Title: ititle, Content: icontent, Pubtime: ipubtime}, nil

}
func CheckErr(e error, statement string) {
	fmt.Println(statement)
	panic(e)
}

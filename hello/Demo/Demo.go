package main

import (
	"fmt"

	. "git.woa.com/qingruixu/Mysql"
	. "git.woa.com/qingruixu/Server"
)

func main() {
	err := InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败 %v", err)
		return
	}

	err = InsertToMysql("Jony", "this is Jony", 70, "Jony")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = RunServer()
	if err != nil {
		fmt.Println("服务器启动失败 %v", err)
		return
	}
	defer DB.Close()

}

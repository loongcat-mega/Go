package Server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "git.woa.com/qingruixu/Mysql"
	//. "git.woa.com/qingruixu/News"
	//"io/ioutil"
	//
)

func getHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	q := r.URL.Query()

	if len(q) != 0 {

		id, err := strconv.Atoi(q.Get("id"))
		if err != nil {
			panic(err)
		}

		res, err := QueryFromMysql(id)
		if err != nil {
			fmt.Println("查询失败 %v", err)
			return
		}
		//po := &res
		bits, err := json.Marshal(res)

		if err != nil {
			fmt.Println("转json失败 %v ", err)
			return
		}

		fmt.Println(string(bits))
		fmt.Fprintln(w, string(bits))
	}
}

func RunServer() error {

	fmt.Println("Server already open")
	http.HandleFunc("/get", getHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		return fmt.Errorf("端口监听失败失败:%v ", err)
	}
	return nil
}

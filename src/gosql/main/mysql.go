package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "frank:5216@tcp(localhost:3306)/go_sql?charset=utf8")
	checkErr(err)

	//插入数据
	fmt.Println("---------插入数据--------")
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("frank", "研发部门", "2018-07-23")
	res, err = stmt.Exec("liuk", "财务部门", "2018-07-21")
	res, err = stmt.Exec("ghost", "人力资源部门", "2018-06-24")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("LastInsertId", id)

	//更新数据
	fmt.Println("---------更新数据-------")
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("frankliu", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("RowsAffected:", affect)

	//查询数据
	rows, err := db.Query("select * from userinfo");
	checkErr(err)

	fmt.Println("-------查询结果:--------")
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println("uid:", uid)
		fmt.Println("usrname:", username)
		fmt.Println("departname:", departname)
		fmt.Println("created", created)
	}

	//删除数据
	fmt.Println("---------删除数据---------")
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("deleteAffected:", affect)

	db.Close()   //关闭数据库连接
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
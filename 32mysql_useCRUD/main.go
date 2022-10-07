package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
// 定义一个全局的DB句柄
var db *sql.DB

//定义一个数据结构接收sql语句结果
type user struct {
	id int
	name string
	age int
}

//initDB 初始化DB：连接数据库
func initDB() (err error) {
	//数据库信息
	//用户名:密码@tcp(ip:port)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Println("dsn",err)
		return err
	}
	//设置数据连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置数据连接池的最大空闲连接数
	db.SetMaxIdleConns(5)
	return nil
}
//queryROW 单行查询
func queryRow() error {
	sqlStr := `select id,name,age from user where id=?;`
	//
	var u user
	//必须调用Scan方法，释放连接
	err := db.QueryRow(sqlStr,2).Scan(&u.id,&u.name,&u.age)
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}

//queryRows 多行查询
func queryRows() error {
	sqlStr := `select id,name,age from user where id>?;`
	rows,err := db.Query(sqlStr,0)
	if err != nil {
		return err
	}
	//非常重要：关闭释放rows持有的数据库连接
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.Scan(&u.id,&u.name,&u.age)
		if err != nil {
			fmt.Printf("rows Scan failed, err:%v\n",err)
			return err
		}
		fmt.Println(u)
	}
	return nil
}

//insertRow 单行插入
func insertRow() error{
	sqlStr := `insert into user(name,age) values(?,?);`
	//执行数据插入语句
	ret,err := db.Exec(sqlStr,"吴凡",24)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n",err)
		return err
	}
	//数据插入结束
	theID, err:=ret.LastInsertId()//新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsertID failed,err:%v\n",err)
		return err
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return nil
}
//updateRow 更新数据
func updateRow() error {
	sqlStr := `update user set age=? where id=?;`
	//执行数据更新语句
	ret,err := db.Exec(sqlStr,16,2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	rowNum, err := ret.RowsAffected()//返回操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n",err)
		return err
	}
	fmt.Printf("update success, affected rows:%v\n",rowNum)
	return nil
}

//deleteRow 删除数据
func deleteRow() error {
	sqlStr := `delete from user where id = ?;`
	ret, err := db.Exec(sqlStr, 5)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n",err)
		return err
	}
	rowNum, err := ret.RowsAffected()//返回操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n",err)
		return err
	}
	fmt.Printf("delete success, affected rows:%v\n",rowNum)
	return nil
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
	//查询表中的单行数据
	err = queryRows()
	if err != nil {
		fmt.Println(err)
	}
	err = insertRow()
	if err != nil {
		fmt.Println(err)
	}
	err = queryRows()
	if err != nil {
		fmt.Println(err)
	}
	err = updateRow()
	if err != nil {
		fmt.Println(err)
	}
	err = queryRows()
	if err != nil {
		fmt.Println(err)
	}
	err = deleteRow()
	if err != nil {
		fmt.Println(err)
	}
	err = queryRows()
	if err != nil {
		fmt.Println(err)
	}

}
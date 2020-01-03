package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "merlin:helloworld@tcp(localhost:3306)/superetail?charset=utf8")
}

func main() {
	//insert()
	query()
	//update()
	//query()
	//delete()
}

func query() {
	tmpSql := "SELECT id,skuName,jdPrice FROM JDSkus WHERE name='食品饮料' AND skuName LIKE '%矿泉水%'"
	//方式1 query
	rows, _ := db.Query(tmpSql)
	defer rows.Close()
	if rows == nil {
		fmt.Println("rows is null!")
		return
	}
	for rows.Next() {
		var id string
		var jdPrice float32
		var skuName string
		if err := rows.Scan(&id, &skuName, &jdPrice); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, skuName:%s ,jdPrice: %f\n", id, skuName, jdPrice)
	}
}

func update() {
	start := time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		tx.Exec("UPdate user set age=? where uid=?", i, i)
	}
	tx.Commit()

	end := time.Now()
	fmt.Println("方式4 update total time:", end.Sub(start).Seconds())
}

func insert() {
	start := time.Now()
	//Begin函数内部会去获取连接
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec("INSERT INTO user(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
	}
	//最后释放tx内部的连接
	tx.Commit()

	end := time.Now()
	fmt.Println("方式4 insert total time:", end.Sub(start).Seconds())

}

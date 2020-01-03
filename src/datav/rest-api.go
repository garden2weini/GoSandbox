package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "merlin:helloworld@tcp(localhost:3306)/superetail?charset=utf8")
}

// TODO: 活跃设备日均销售额
func getTODO(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling getTODO handler")
	w.Header().Set("Content-Type", "application/json")
	var datas = []Data4Axis{}

	tmpSql := "SELECT createdDate, orders, url FROM Ad LIMIT 100"
	rows, _ := db.Query(tmpSql)
	defer rows.Close()
	if rows == nil {
		fmt.Println("rows is null!")
		return
	}
	for rows.Next() {
		var date string
		var orders int
		var url string
		if err := rows.Scan(&date, &orders, &url); err != nil {
			log.Fatal(err)
		}
		data1 := Data4Axis{X: date, Y: strconv.Itoa(orders), Z: url}
		datas = append(datas, data1)
		//fmt.Printf("Date:%s ,Orders: %d, Url: %s\n", date, orders, url)
	}

	json.NewEncoder(w).Encode(datas)
}

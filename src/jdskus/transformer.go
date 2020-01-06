package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_NAME = "merlin"
	PASS_WORD = "helloworld"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "superetail"
	CHARSET   = "utf8"
	//PARAMS    = "useUnicode=true&characterEncoding=UTF-8&serverTimezone=GMT"
)

var db = &sql.DB{}

func init() {
	// parseTime=true 可以解析dateTime
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	fmt.Println(dbDSN)
	db, _ = sql.Open("mysql", "merlin:helloworld@tcp(localhost:3306)/superetail?charset=utf8")
	//db, _ = sql.Open("mysql", dbDSN)
}

func main() {
	transerProduct()
}

func transerProduct() {
	tmpSql := "SELECT id,skuName,jdPrice FROM JDSkus WHERE name='食品饮料' AND skuName LIKE '%矿泉水%' ORDER BY skuName"

	rows, _ := db.Query(tmpSql)
	defer rows.Close()
	if rows == nil {
		fmt.Println("rows is null!")
		return
	}
	for rows.Next() {
		var jdSkuId int
		pro1 := NewProduct()
		sku1 := NewSku()
		sku1.createdDate = time.Now()

		if err := rows.Scan(&jdSkuId, &pro1.name, &pro1.price); err != nil {
			log.Fatal(err)
			return
		}
		kv := strings.Split(pro1.name, "*")
		if len(kv) != 2 {
			fmt.Printf("Ops：" + pro1.name + "\n")
			return
		} else {
			pro1.name = kv[0]

		}
		//fmt.Printf("ID:%d, skuName:%s, jdPrice:%f\n", jdSkuId, pro1.name, pro1.price)
		pro1.id = nextTableId("Product")
		insertProduct(pro1)
	}
}

// 将JDSku与Product关联
func updateJDSku(jdSkuId int, productId int) {
	tx, _ := db.Begin()
	tx.Exec("Update JDSkus set product_id=? where id=?", productId, jdSkuId)
	tx.Commit()

}

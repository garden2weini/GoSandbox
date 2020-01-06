package main

import (
	"database/sql"
	"fmt"
	"strings"

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
	jdSkuList := getJDSkuList("食品饮料", "矿泉水")
	//var skuList []Sku
	var productMap map[string]Product

	for index, jdSku := range jdSkuList {
		fmt.Printf("index slice[%d] = %s\n", index, jdSku.skuName)
		// jdSku.skuName期望分为名称和数量+单位等两部分
		kv := strings.Split(jdSku.skuName, "*")
		if len(kv) != 2 {
			fmt.Printf("Ops：" + jdSku.skuName + "\n")
			continue
		} else {
			var product Product
			if v, ok := productMap[kv[0]]; ok == true {
				// 对应Product不存在
				product = NewProduct()
				sku := NewSku()
				product.id = nextTableId("Product")
				sku.id = nextTableId("Sku")
				product.name = kv[0]
				sku.price = jdSku.jdPrice
				sku.product_id = product.id
				//skuList = append(skuList, sku)
				productMap[product.name] = product
				insertProduct(product)
				insertSku(sku)
			} else {
				// 对应Product已经存在
				product = v
			}
			jdSku.product_id = product.id
			updateJDSku(jdSku)
		}

	}

}

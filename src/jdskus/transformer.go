package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
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
	// 声明并初始化map
	productMap := map[string]Product{}

	for index, jdSku := range jdSkuList {
		fmt.Printf("index slice[%d] = %s\n", index, jdSku.skuName)
		// jdSku.skuName期望分为名称和数量+单位等两部分（如：大傻子矿泉水300ml*24瓶）
		kv := strings.Split(jdSku.skuName, "*")
		if len(kv) != 2 {
			fmt.Printf("Ops：" + jdSku.skuName + "\n")
			continue
		} else {
			var product Product
			if v, ok := productMap[kv[0]]; ok == false {
				// 对应Product不存在
				product = NewProduct()
				sku := NewSku()
				jdSku.quantity = parseQuantity(kv[1])
				product.id = nextTableId("Product")
				product.name = kv[0]
				product.price = jdSku.jdPrice / float32(jdSku.quantity)
				sku.id = nextTableId("Sku")
				sku.price = jdSku.jdPrice
				sku.product_id = product.id
				//fmt.Printf("jdSku.productCategory_id：%d\n", jdSku.productCategory_id)
				if jdSku.productCategory_id != 0 {
					product.productCategory_id = jdSku.productCategory_id
				}
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

func parseQuantity(rawStr string) int {
	quantity := 1
	pattern := "^[1-9]\\d*$" //反斜杠要转义
	for i := 1; i < 10; i++ {
		subStr := rawStr[0:i]
		result, _ := regexp.MatchString(pattern, subStr)
		//fmt.Println(result)
		if !result {
			break
		} else {
			//fmt.Println(subStr)
			quantity, _ = strconv.Atoi(subStr)
		}
	}
	return quantity
}

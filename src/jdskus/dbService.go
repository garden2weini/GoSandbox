package main

import (
	"fmt"
	"log"
)

func getJDSkuList(skuType string, filter string) []JDSku {
	var jdSkuSlice []JDSku
	tmpSql := "SELECT id,skuName,jdPrice FROM JDSkus WHERE name=? AND skuName LIKE ? ORDER BY skuName"
	//selectSql := fmt.Sprintf(tmp2, tableName)
	rows, err := db.Query(tmpSql, skuType, "%"+filter+"%")
	if err != nil {
		log.Fatal(err)
		return jdSkuSlice
	}

	defer rows.Close()
	if rows == nil {
		fmt.Println("rows is null!")
		return jdSkuSlice
	}
	for rows.Next() {
		jdSku := NewJDSku()

		if err := rows.Scan(&jdSku.id, &jdSku.skuName, &jdSku.jdPrice); err != nil {
			log.Fatal(err)
			continue
		}
		jdSkuSlice = append(jdSkuSlice, jdSku)
	}
	return jdSkuSlice
}

// 插入Sku表
func insertSku(item Sku) int64 {
	skuParams := [...]string{"id", "createdDate", "lastModifiedDate", "version",
		"allocatedStock", "exchangePoint", "isDefault", "marketPrice",
		"maxCommission", "price", "rewardPoint", "sn", "stock", "product_id"}
	sql := "INSERT INTO Sku(%s) values(%s)"

	sql = buildSqlByColumns(skuParams[:], sql)
	//fmt.Println(sql)

	//Begin函数内部会去获取连接
	tx, _ := db.Begin()
	result, err := db.Exec(sql, item.id, item.createdDate, item.lastModifiedDate, item.version,
		item.allocatedStock, item.exchangePoint, item.isDefault, item.marketPrice,
		item.maxCommission, item.price, item.rewardPoint, item.sn, item.stock, item.product_id)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	//最后释放tx内部的连接
	tx.Commit()
	rowCnt, _ := result.RowsAffected()
	return rowCnt
}

// 插入Product表
func insertProduct(item Product) int64 {
	productParams := [...]string{"id", "createdDate", "lastModifiedDate", "version",
		"hits", "isActive", "isDelivery", "isList", "isMarketable", "isTop",
		"marketPrice", "maxCommission", "monthHits", "monthHitsDate", "monthSales",
		"monthSalesDate", "name", "price", "sales", "score", "scoreCount", "sn", "totalScore",
		"type", "weekHits", "weekHitsDate", "weekSales", "weekSalesDate",
		"productCategory_id"}
	sql := "INSERT INTO Product(%s) values(%s)"

	sql = buildSqlByColumns(productParams[:], sql)
	//fmt.Println(sql)

	//Begin函数内部会去获取连接
	tx, _ := db.Begin()
	result, err := db.Exec(sql, item.id, item.createdDate, item.lastModifiedDate,
		item.version, item.hits, item.isActive, item.isDelivery, item.isList,
		item.isMarketable, item.isTop, item.marketPrice, item.maxCommission,
		item.monthHits, item.monthHitsDate, item.monthSales, item.monthSalesDate,
		item.name, item.price, item.sales, item.score, item.scoreCount, item.sn, item.totalScore,
		item._type, item.weekHits, item.weekHitsDate, item.weekSales, item.weekSalesDate,
		item.productCategory_id)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	//最后释放tx内部的连接
	tx.Commit()
	rowCnt, _ := result.RowsAffected()
	return rowCnt
}

// 将JDSku与Product关联及Product数量
func updateJDSku(jdSku JDSku) {
	tx, _ := db.Begin()
	_, err := tx.Exec("Update JDSkus set product_id=?, quantity=? where id=?", jdSku.product_id, jdSku.quantity, jdSku.id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

}

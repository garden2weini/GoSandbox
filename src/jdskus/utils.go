package main

import (
	"fmt"
	"log"
	"strings"
)

func buildSqlByColumns(params []string, sqlTemplate string) string {
	var p1 strings.Builder
	var p2 strings.Builder
	for i := 0; i < len(params); i++ {
		if i == len(params)-1 {
			p1.WriteString(params[i])
			p2.WriteString("?")
		} else {
			p1.WriteString(params[i])
			p1.WriteString(",")
			p2.WriteString("?,")
		}
	}
	//sql := "INSERT INTO Product(%s) values(%s)"
	sql := fmt.Sprintf(sqlTemplate, p1.String(), p2.String())
	return sql
}

// 获取指定Table的nextVal(通过Table IdGenerator[sequence_name,next_val])
func nextTableId(tableName string) int {
	updateSql := "update IdGenerator set next_val=next_val+1 where sequence_name=?"
	tmp2 := "select next_val from IdGenerator where sequence_name='%s'"
	// Sprintf格式化并返回一个字符串而不带任何输出
	selectSql := fmt.Sprintf(tmp2, tableName)
	//fmt.Println(updateSql)
	//fmt.Println(selectSql)

	stm, _ := db.Prepare(updateSql)
	_, err := stm.Exec(tableName)
	if err != nil {
		fmt.Println("Fail to update nextVal!")
		log.Fatal(err)
		return -1
	}
	stm.Close()

	rows, err := db.Query(selectSql)
	if err != nil {
		fmt.Println("Fail to exec nextVal query!")
		log.Fatal(err)
		return -1
	}
	defer rows.Close()
	if rows == nil {
		fmt.Println("IdGenerator rows is null!")
		return -1
	}
	for rows.Next() {
		var nextVal int

		if err := rows.Scan(&nextVal); err != nil {
			fmt.Println("Fail to get nextVal!")
			log.Fatal(err)
			return -1
		}
		return nextVal
	}
	return -1
}

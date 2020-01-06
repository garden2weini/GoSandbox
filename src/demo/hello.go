package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	testMap := map[string]int{}
	testMap["1"] = 1
	fmt.Println(testMap)

	count := 1
	pattern := "^[1-9]\\d*$" //反斜杠要转义
	str := "124瓶"
	result, _ := regexp.MatchString(pattern, str)
	fmt.Println(result)
	for i := 1; ; i++ {
		subStr := str[0:i]
		result, _ := regexp.MatchString(pattern, subStr)
		fmt.Println(result)
		if !result {
			fmt.Printf("Fail:%s", subStr)
			break
		} else {
			fmt.Println(subStr)
			count, _ = strconv.Atoi(subStr)
		}
	}
	fmt.Printf("Count:%d", count)

}

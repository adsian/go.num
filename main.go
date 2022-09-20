package main

import (
	"fmt"
	"github.com/Tnze/go.num/v2/zh"
)

func main() {
	var num uint64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(zh.Uint64(num).String())
	fmt.Println(zh.Uint64(num).StringTraditional())

	var num2 zh.Uint64
	num3, _ := num2.Scan("四十五万七千四百三十五")
	fmt.Println(num3)
}

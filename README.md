# go.num

Mainly implemented int->Chinese
主要实现了整数转中文算法

```go
package main

import (
	"fmt"
	"github.com/adsian/go.num/v2/zh"
)

func main() {
	// 数字转中文
	var num uint64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(zh.Uint64(num).String())
	// 转繁体中文
	fmt.Println(zh.Uint64(num).StringTraditional())
}
```

```go
func main() {
	// 中文转数字
	var num2 zh.Uint64
	num3, _ := num2.Scan("四十五万七千四百三十五")
    fmt.Println(num3)
}
```

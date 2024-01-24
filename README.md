## 示例代码
```
package main

import (
	"github.com/kimistar/go-ini"
	"fmt"
)

func main() {
	filename := "testdata/more.ini"
	section := "development"

	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg.Read(section, "name"))
	fmt.Println(cfg.Read(section, "name1"))
}

```
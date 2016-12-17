# dogo

dogo 用go语言开发的web框架，高效，简单。

```
package main

import (
	"github.com/wuciyou/dogo"
	_ "github.com/wuciyou/dogo/example/web/controller"
)

func main() {
	dogo.Start()
}

```
package main

import (
	_ "gfim/boot"
	_ "gfim/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

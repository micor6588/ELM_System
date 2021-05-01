package main

import (
	_ "ELM_System/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

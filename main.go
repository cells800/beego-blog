package main

import (
	"github.com/astaxie/beego"
	_ "igo9go/beego-blog/models"
	_ "igo9go/beego-blog/routers"
)

func main() {
	beego.Run()
}

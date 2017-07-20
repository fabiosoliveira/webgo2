package main

import (
	"github.com/astaxie/beego"
	_ "github.com/fabiosoliveira/webgo/routers"
	"os"
	"strconv"
	"log"
)

func main() {

	if porta, erro := strconv.Atoi(os.Getenv("PORT")) ; erro == nil {
		beego.BConfig.Listen.HTTPPort = porta
	} else {
		log.Fatalf("Erro ao passar o http port como argumento: %v", erro.Error())
	}

	beego.Run()
}

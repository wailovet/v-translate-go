package main

import (
	"flag"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"github.com/wailovet/v-translate-go/translate"
)

func main() {

	var port string
	flag.StringVar(&port, "port", "20080", "端口号:默认[20080]")
	flag.Parse()

	core.GetInstanceConfig().Port = port
	core.GetInstanceConfig().UpdatePath = "v_translate_linux_update"
	core.GetInstanceConfig().StaticRouter = "/T/*"

	core.GetInstanceRouterManage().Registered(&translate.Index{})

	osmanthuswine.Run()
}

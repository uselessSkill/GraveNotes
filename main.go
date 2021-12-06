package main

import (
	"studyGin/conf"
	"studyGin/routers"
)

func main() {
	config := conf.Init()
	r := routers.Init()
	_ = r.Run(config.Http.Listen)
}

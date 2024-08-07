package main

import (
	"gin-vue-blog/model"
	"gin-vue-blog/routers"
)

func main() {
	model.InitDB()
	routers.InitRouter()
}

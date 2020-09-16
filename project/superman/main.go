package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shuai/superman/manage"
)

func Init() {
	manage.RedisInit()
	manage.InitRedisPool()

}

func main() {
	Init()
	//manage.InitWatcher(context.Background())
	router := gin.Default()
	RegisterRouter(router)
	http.ListenAndServe(":8088", router)

}

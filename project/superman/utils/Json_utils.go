package utils

import (
	"github.com/gin-gonic/gin"
	"shuai/superman/consts"
)

func JSONSucc(ctx *gin.Context, status *consts.Status, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": status.Code,
		"msg":  status.Msg,
		"data": data,
	})
}

func JSONFail(ctx *gin.Context, status *consts.Status) {
	ctx.JSON(200, gin.H{
		"code": status.Code,
		"msg":  status.Msg,
		"data": gin.H{},
	})
}

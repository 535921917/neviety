package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shuai/superman/model"
)

const DefaultTTL = 10

func CheckGinServiceItem(ctx *gin.Context, item *model.ServiceItem) error {
	if err := ctx.ShouldBindQuery(item); err != nil {
		log.Printf("RegisterOrKeep ShouldBindJSON :err = %v", err)
		return err
	}
	if item.PSM == "" || item.IP == "" || item.Port == 0 {
		log.Printf("illegal item = %+v", item)
		return fmt.Errorf("illegal param error")
	}
	if item.TTL <= 0 {
		item.TTL = DefaultTTL
	}
	return nil
}

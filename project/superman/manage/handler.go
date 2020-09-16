package manage

import (
	"github.com/gin-gonic/gin"
	"shuai/superman/consts"
	"shuai/superman/model"
	"shuai/superman/utils"
)

func RegisterOrKeep(ctx *gin.Context) {
	serviceItem := &model.ServiceItem{}
	if err := CheckGinServiceItem(ctx, serviceItem); err != nil {
		utils.JSONFail(ctx, consts.ERR_ILLEGAL_PARAM)
		return
	}
	ServiceKeep(ctx, serviceItem)
	utils.JSONSucc(ctx, consts.OK, nil)
}
func Cancel(ctx *gin.Context) {
	serviceItem := &model.ServiceItem{}
	if err := CheckGinServiceItem(ctx, serviceItem); err != nil {
		utils.JSONFail(ctx, consts.ERR_ILLEGAL_PARAM)
		return
	}
	CancelKeep(ctx, serviceItem)
	utils.JSONSucc(ctx, consts.OK, nil)
}

func GetServiceList(ctx *gin.Context) {
	psmArray := ctx.QueryArray("PSM")
	if len(psmArray) == 0 {
		utils.JSONFail(ctx, consts.ERR_ILLEGAL_PARAM)
		return
	}
	list := ServiceList(ctx, psmArray)
	utils.JSONSucc(ctx, consts.OK, list)
}

func GetServiceVersion(ctx *gin.Context) {
	psmArray := ctx.QueryArray("PSM")
	if len(psmArray) == 0 {
		utils.JSONFail(ctx, consts.ERR_ILLEGAL_PARAM)
		return
	}
	m := ServiceVersion(ctx, psmArray)
	utils.JSONSucc(ctx, consts.OK, m)
}

/*func StopWatch(ctx *gin.Context) {
  	var cancelFunc context.CancelFunc
  	ctx, cancelFunc = context.WithCancel(ctx)
  	cancelFunc()
  }
*/

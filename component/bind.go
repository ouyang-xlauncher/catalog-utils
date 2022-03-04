package component

import (
	"github.com/gin-gonic/gin"
)

type Binder func(ctx *gin.Context, obj interface{}) error

var (
	defaultBinders = []Binder{BindJSON, BindQuery, BindHeader, BindUri}
)

//Bind 用户请求参数绑定
func Bind(ctx *gin.Context, object *Notify, binders ...Binder) error {
	if len(binders) == 0 {
		binders = defaultBinders
	}
	for _, binder := range binders {
		if binder != nil {
			err := binder(ctx, object)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

//BindJSON 用于绑定请求体
func BindJSON(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		if err.Error() != "EOF" {
			return err
		}
	}
	return nil
}

//BindHeader 用于绑定请求头
func BindHeader(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindHeader(obj)
}

//BindQuery 用于绑定query参数
func BindQuery(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindQuery(obj)
}

//BindUri 用于绑定路径参数
func BindUri(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindUri(obj)
}

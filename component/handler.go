package component

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Interface interface {
	// Add 添加组件
	Add(arg AddArg) error
	// Edit 修改组件
	Edit(arg EditArg) error
	// Del 删除组件
	Del(arg DeleteArg) error
}

type Notify struct {
	*gin.Context
	AddHandler
	EditHandler
	DeleteHandler
	AddArg
	EditArg
	DeleteArg
}

type AddHandler func(AddArg) error
type EditHandler func(EditArg) error
type DeleteHandler func(DeleteArg) error

func NewComponentHandler(ctx *gin.Context, catalogInterface Interface) {
	c := &Notify{
		Context:       ctx,
		AddHandler:    catalogInterface.Add,
		EditHandler:   catalogInterface.Edit,
		DeleteHandler: catalogInterface.Del,
	}
	err := ctx.ShouldBindHeader(catalogInterface)
	if err != nil {
		c.JSON(400, Fail(err))
		return
	}
	err = ctx.ShouldBindJSON(c)
	if err != nil {
		c.JSON(400, Fail(err))
		return
	}
	switch ctx.Request.Method {
	case http.MethodPost:
		err = c.AddHandler(c.AddArg)
	case http.MethodPut:
		err = c.EditHandler(c.EditArg)
	case http.MethodDelete:
		err = c.DeleteHandler(c.DeleteArg)
	}
	if err != nil {
		c.JSON(400, Fail(err))
		return
	}

	c.JSON(200, Success())
	return
}

type Res struct {
	Status int    `json:"status"`
	ResMsg string `json:"resMsg"`
}

func Success() Res {
	return Res{Status: 200}
}

func Fail(err error) Res {
	return Res{Status: 200, ResMsg: err.Error()}
}

package models

import (
	"github.com/astaxie/beego/context"
)

type Context struct {
	Input  *context.BeegoInput
	Output *context.BeegoOutput
	Ctx    *context.Context
}

var ctx Context

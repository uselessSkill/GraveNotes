package routers

import (
	"github.com/gin-gonic/gin"
	_ "studyGin/docs"
	base "studyGin/routers/base"
)

type Option func(*gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}

func init() {
	Include(base.NewsRouters)
}

// 路由注册
func Init() *gin.Engine {
	g := gin.Default()
	for _, opt := range options {
		opt(g)
	}
	return g
}

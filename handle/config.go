package handle

import (
	"github.com/gin-gonic/gin"
	"studyGin/conf"
	"studyGin/output"
)

func GetConfig(c *gin.Context) {
	cf := conf.Init()
	output.Json(c, output.Success, cf)
}

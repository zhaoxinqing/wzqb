package controller

import (
	"Kilroy/app/common"

	"github.com/gin-gonic/gin"
)

func TestGet(c *gin.Context) {
	common.ResSuccess(c, "getAdmin")
}

func TestPost(c *gin.Context) {
	common.ResSuccess(c, "postAdmin")
}
func TestPut(c *gin.Context) {
	common.ResSuccess(c, "putAdmin")
}
func TestDelete(c *gin.Context) {
	common.ResSuccess(c, "deleteAdmin")
}

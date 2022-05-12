package wzqb

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AllMenu 获取所有菜单
func GetAllMenu(c *gin.Context) {
	c.String(http.StatusOK, "The available groups are [...]")
}

// LoginParam ...
type LoginParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// AllMenu 获取所有菜单
func Login(c *gin.Context) {
	resp := LoginParam{
		UserName: "wzqb",
		Password: "liao",
	}

	c.JSON(http.StatusOK, resp)
}

// ArrayToString ...
func ArrayToString(c *gin.Context) {
	wzList := []string{"aaaaa", "sssss", "hhhhh", "ujjj"}
	resp := strings.Join(wzList, ",")
	c.String(http.StatusOK, resp)
}

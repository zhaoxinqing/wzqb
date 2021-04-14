package middlewares

import "github.com/gin-gonic/gin"

// UserAuthMiddleware 用户授权中间件
func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			method = c.Request.Method
			// url    = c.Request.URL.Path
			// token         = c.GetHeader(common.TOKEN_KEY)
		)
		if method != "GET" {
			switch method {
			// case "POST":
			// 	switch url {
			// 	case "/api/admins":
			// 		OperationName = "创建新用户"
			// 		Module = "用户管理模块"
			// 	case "/api/groups":
			// 		OperationName = "创建用户组"
			// 		Module = "用户管理模块"
			// 	case "/api/roles":
			// 		OperationName = "创建新角色"
			// 		Module = "角色管理模块"
			// 	}
			// case "DELETE":
			// 	switch url {
			// 	case "/api/admins":
			// 		OperationName = "删除用户"
			// 		Module = "用户管理模块"
			// 	case "/api/group/":
			// 		OperationName = "删除用户组"
			// 		Module = "用户管理模块"
			// 	case "/api/roles":
			// 		OperationName = "删除角色"
			// 		Module = "角色管理模块"
			// 	}
			// case "PUT":
			// 	switch url {
			// 	case "/api/groups":
			// 		OperationName = "更新用户组"
			// 		Module = "用户管理模块"
			// 	case "/api/roles":
			// 		OperationName = "更新角色"
			// 		Module = "角色管理模块"
			// 	case "/api/admins":
			// 		OperationName = "更新用户"
			// 		Module = "用户管理模块"
			// 	}
			// }
			// userName, _ := service.GetUserName(userID, token)
			// var logParam = OperationLogParam{
			// 	tenantID,
			// 	userName,
			// 	method,
			// 	OperationName,
			// 	Module,
			// 	c.ClientIP(),
			// }
			// _ = AddOperationLog(logParam)
			}
		}
	}
}

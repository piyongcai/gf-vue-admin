package system

import (
	v1 "server/app/api/v1"
	"server/library/global"
)

// InitBaseRouter 注册基础功能路由
func InitBaseRouter() {
	BaseRouter := global.GFVA_SERVER.Group("base")
	{
		BaseRouter.POST("register", v1.Register)                  // 注册
		BaseRouter.POST("login", v1.GfJWTMiddleware.LoginHandler) // 登录
		BaseRouter.POST("captcha", v1.Captcha)                    // 验证码
	}
}
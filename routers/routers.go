package routers

import (
	"github.com/gin-gonic/gin"
	"go-session-demo/controllers"
	"go-session-demo/helpers"
	"go-session-demo/models/response"
)

/*
*
需要登录鉴权的中间件
*/
func TokenRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Token")
		if len(token) == 0 {
			ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
				MsgCode: helpers.MSG_CODE_NOT_LOGIN,
				Desc:    helpers.MSG_DESC_NOT_LOGIN,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func RouterRun() {
	routers := gin.New()

	apiRouters := routers.Group("/")

	apiRouters.Group("/").Use(gin.Logger()).Use(gin.Recovery())
	{
		apiRouter := routers.Group("/api")
		apiV1 := apiRouter.Group("/v1")

		// account路由组
		accountRouter := apiV1.Group("/account")
		account := &controllers.Account{}

		// 不需要token检测
		accountRouter.GET("/register", account.Register)
		accountRouter.GET("/login", account.Login)

		accountRouter.Use(TokenRequired())
		{
			accountRouter.GET("/list", account.List)
			accountRouter.GET("/edit", account.UpdInfo)
			accountRouter.GET("/logout", account.Logout)
			accountRouter.GET("/check", account.Check)
			accountRouter.GET("/close", account.Close)
		}
	}

	_ = routers.Run()
}

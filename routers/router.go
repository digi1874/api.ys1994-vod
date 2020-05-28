/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 15:50:01
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-25 11:15:22
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-vod/controllers"
	"api.ys1994-vod/process"
)

// Run run Router
func Run() {
	if process.IsDev == false {
		gin.SetMode(gin.ReleaseMode)
	}

	var Router = gin.Default()
	Router.Use(middleware())

	vod := Router.Group("/vod")

	vod.GET("/list", controllers.GetVodListHandle)
	vod.GET("/detail/:id", controllers.GetVodDetailHandle)
	vod.GET("/keyword", controllers.GetKeywordListHandle)

	if process.IsDev == false {
		Router.Run("127.0.0.1:8050")
	} else {
		Router.Run("127.0.0.1:8051")
	}
}

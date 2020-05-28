/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 10:58:17
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-20 11:05:42
 */

package main

import (
	"api.ys1994-vod/database"
	"api.ys1994-vod/routers"
)

func main()  {
	defer database.DB.Close()
	routers.Run()
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 11:11:22
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-24 20:56:16
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"api.ys1994-vod/model"
)

// GetKeywordListHandle 关键字列表
func GetKeywordListHandle(c *gin.Context) {
	var err error
	var vkl model.VodKeywordList
	vkl.Page, vkl.Size, err = listHandle(c, &vkl.Filter, &vkl.Filters)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}
	vkl.Find()
	cJSONOk(c, vkl)
}

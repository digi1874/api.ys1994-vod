/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 16:13:22
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-25 17:23:08
 */

package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// cJSONBadRequest 错误请求返回
func cJSONBadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{ "msg": msg })
}

// cJSONNotFound 404请求返回
func cJSONNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{ "msg": msg })
}

// cJSONOk 正常请求返回
func cJSONOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{ "data": data })
}

// cJSONUnauthorized 没权限
func cJSONUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{ "msg": msg })
}
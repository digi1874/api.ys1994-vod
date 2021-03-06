/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 11:38:08
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-20 18:51:22
 */

package controllers

import (
	"strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func listHandle(c *gin.Context, filter1 interface{}, filter2 interface{}) (page int, size int, err error) {
	page, _ = strconv.Atoi(c.Query("page"))
	size, _ = strconv.Atoi(c.Query("size"))
	ft      := c.Query("filter")

	if ft != "" {
		err = json.Unmarshal([]byte(ft), &filter1)

		if err == nil {
			err = json.Unmarshal([]byte(ft), &filter2)
		}
	}

	return page, size, err
}

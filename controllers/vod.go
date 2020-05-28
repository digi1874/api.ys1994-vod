/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 20:47:58
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-25 17:23:23
 */

package controllers

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"api.ys1994-vod/model"
)

var serialRe = regexp.MustCompile(`更新`)

// GetVodDetailHandle  Vod detail
func GetVodDetailHandle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 1 {
		cJSONBadRequest(c, "ID错误，ID只为正整数")
		return
	}

	vod := model.NewVod()
	vod.Filter.ID = uint(id)
	vod.Filter.State = 1
	vod.Detail()

	if vod.ID == 0 {
		cJSONNotFound(c, "视频不存在")
		return
	}

	if serialRe.MatchString(vod.Serial) {
		if vod.UpdatedTime + 3600 < uint(time.Now().Unix()) && search(vod.Name, c.Query("sync") != "") {
			v := model.NewVod()
			v.Filter.ID = vod.ID
			v.Detail()
			vod = v
		}
	}

	cJSONOk(c, vod)
}

// GetVodListHandle vod列表
func GetVodListHandle(c *gin.Context) {
	var err error
	var vodList model.VodList
	vodList.Page, vodList.Size, err = listHandle(c, &vodList.Filter, &vodList.Filters)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	if vodList.Filter.Name != "" {
		search(vodList.Filter.Name, c.Query("sync") != "")
	}

	if vodList.Filter.State != 2 {
		// 留后路查禁用影片
		vodList.Filter.State = 1
	}

	vodList.Find()
	cJSONOk(c, vodList)
}

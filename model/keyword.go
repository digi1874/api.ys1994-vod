/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 18:13:10
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-24 20:09:16
 */

package model

import (
	"api.ys1994-vod/database"
)

// VodKeyword 关键字
type VodKeyword struct {
	Base
	ID               uint                 `json:"-"`
	Text             string               `json:"text"`
	Num              uint                 `json:"num"`
	SyncTime         uint                 `json:"-"`
	SyncLoading      uint8                `json:"-"`
	Filter           database.VodKeyword  `json:"-"`
}

// NewVodKeyword New VodKeyword
func NewVodKeyword() *VodKeyword {
	var vk VodKeyword
	vk.Base.Super(&vk, &vk.Filter)
	return &vk
}

// VodKeywordList 关键字列表
type VodKeywordList struct {
	List
	Data             []VodKeyword        `json:"data"`
	Filter           database.VodKeyword `json:"-"`
}

// IncrementNum 增加数量；VodKeyword.Filter.ID != 0
func (vk *VodKeyword) IncrementNum() {
	vk.Base.Increment("num")
}

// Find 列表
func (vkl *VodKeywordList) Find() {
	// 不能查删除的
	vkl.Filter.DeletedAt = nil
	db := database.DB.Where(&vkl.Filter)

	vkl.dbHandle(db.Order("num desc").Order("updated_time desc")).Find(&vkl.Data)

	db.Model(&database.VodKeyword{}).Count(&vkl.Count)
}

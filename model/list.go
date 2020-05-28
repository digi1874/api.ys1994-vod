/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-04 17:45:55
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-16 22:29:48
 */

package model

import (
	"github.com/jinzhu/gorm"
)

// Filters 过滤
type Filters struct {
	IDs              []uint   `json:"ids"`
	CreatedTimeStart uint     `json:"createdTimeStart"`
	CreatedTimeEnd   uint     `json:"createdTimeEnd"`
	UpdatedTimeStart uint     `json:"updatedTimeStart"`
	UpdatedTimeEnd   uint     `json:"updatedTimeEnd"`
}

// List 列表
type List struct {
	Count           int        `json:"count"`
	Page            int        `json:"page"`
	Size            int        `json:"size"`
	Filters         Filters    `json:"-"`
}

func (f Filters) dbHandle(db *gorm.DB) *gorm.DB {
	if len(f.IDs) > 0 {
		db = db.Where("`id` IN (?)", f.IDs)
	}
	if f.CreatedTimeStart > 0 {
		db = db.Where("`created_time` >= ?", f.CreatedTimeStart)
	}
	if f.CreatedTimeEnd > 0 {
		db = db.Where("`created_time` <= ?", f.CreatedTimeEnd)
	}
	if f.UpdatedTimeStart > 0 {
		db = db.Where("`updated_time` >= ?", f.UpdatedTimeStart)
	}
	if f.UpdatedTimeEnd > 0 {
		db = db.Where("`updated_time` <= ?", f.UpdatedTimeEnd)
	}
	return db
}

func (l *List) dbHandle(db *gorm.DB) *gorm.DB {
	if l.Page < 1 {
		l.Page = 1
	}
	if l.Size < 1 {
		l.Size = 20
	}
	offset := (l.Page - 1) * l.Size
	return db.Limit(l.Size).Offset(offset)
}

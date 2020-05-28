/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 20:53:17
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-25 21:15:24
 */

package model

import (
	"github.com/jinzhu/gorm"

	"api.ys1994-vod/database"
)

// VOD 用在VodList
type VOD struct {
	Base
	ID              uint               `json:"id"`
	Actor           string             `json:"actor"`
	Area            string             `json:"area"`
	Serial          string             `json:"serial"`
	Director        string             `json:"director"`
	Name            string             `json:"name"`
	Pic             string             `json:"pic"`
	Year            uint16             `json:"year"`
	TypeID          uint8              `json:"typeId"`
}

// Vod 影片
type Vod struct {
	VOD
	SubName         string             `json:"subName"`
	Lang            string             `json:"lang"`
	Content         string             `json:"content"`
	TypePID         uint8              `json:"typePId"`
	UpdatedTime     uint               `json:"updatedTime"`
	URLs            []VodURL           `json:"urls"`
	M3u8s           []VodM3u8          `json:"m3u8s"`
	DownURLs        []VodDownURL       `json:"downURLs"`
	Filter          database.Vod       `json:"-"`
}

// VodList 影片列表
type VodList struct {
	List
	Data             []VOD        `json:"data"`
	Filter           database.Vod `json:"-"`
}

// NewVod New Vod
func NewVod() *Vod {
	var v Vod
	v.Base.Super(&v, &v.Filter)
	return &v
}

// Detail Vod detail
func (v *Vod) Detail() {
	v.Base.Detail()
	if v.ID != 0 {
		v.Base.Related(&v.URLs)
		v.Base.Related(&v.M3u8s)
		v.Base.Related(&v.DownURLs)
	}
}

// Find 列表
func (vl *VodList) Find() {
	vl.Filter.DeletedAt = nil
	db := handleFilterVod(database.DB, vl.Filter)
	db = vl.Filters.dbHandle(db)
	vl.dbHandle(db.Order("updated_time desc")).Find(&vl.Data)
	db.Model(&database.Vod{}).Count(&vl.Count)
}

func handleFilterVod(db *gorm.DB, filter database.Vod) *gorm.DB {
	if filter.Name != "" {
		name := "%" + filter.Name + "%"
		db = db.Where("`vods`.`name` LIKE ? OR `vods`.`sub_name` LIKE ?", name, name)
		filter.Name = ""
	}

	if filter.Actor != "" {
		db = db.Where("`vods`.`actor` LIKE ?", "%" + filter.Actor + "%")
		filter.Actor = ""
	}

	if filter.Director != "" {
		db = db.Where("`vods`.`director` LIKE ?", "%" + filter.Director + "%")
		filter.Director = ""
	}

	if filter.Serial != "" {
		db = db.Where("`vods`.`serial` LIKE ?", "%" + filter.Serial + "%")
		filter.Serial = ""
	}

	if filter.Area != "" {
		db = db.Where("`vods`.`area` LIKE ?", "%" + filter.Area + "%")
		filter.Area = ""
	}

	if filter.Lang != "" {
		db = db.Where("`vods`.`lang` LIKE ?", "%" + filter.Lang + "%")
		filter.Lang = ""
	}

	return db.Where(&filter)
}

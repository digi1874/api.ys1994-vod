/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 11:10:08
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-25 11:32:11
 */

package database

// VodKeyword 关键字
type VodKeyword struct {
	Model
	Text             string    `gorm:"not null;comment:'关键字'"`
	Num              uint      `gorm:"DEFAULT:0;comment:'搜索次数'"`
	SyncTime         uint      `gorm:"DEFAULT:0;comment:'同步资源时间'"`
	SyncLoading      uint8     `gorm:"DEFAULT:1;comment:'同步中；1：没在同步；2: 同步中'"`
}

// Vod 影片结构体
type Vod struct {
	Model
	Oid             uint        `gorm:"comment:'okZy id'"`
	TypeID          uint8       `gorm:"comment:'类型id'"`
	TypePID         uint8       `gorm:"comment:'类型pid'"`
	Name            string      `gorm:"comment:'视频名'"`
	SubName         string      `gorm:"comment:'视频别名'"`
	PY              string      `gorm:"comment:'视频名拼音'"`
	Pic             string      `gorm:"comment:'封面'"`
	Actor           string      `gorm:"comment:'演员'"`
	Director        string      `gorm:"comment:'导演'"`
	Serial          string      `gorm:"comment:'最近更新'"`
	Area            string      `gorm:"comment:'地区'"`
	Lang            string      `gorm:"comment:'语言'"`
	Year            uint16      `gorm:"comment:'年'"`
	Content         string      `gorm:"type:varchar(5000);comment:'绍介'"`
	State           uint8       `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
	URLs            []VodURL
	M3u8s           []VodM3u8
	DownURLs        []VodDownURL
}

// VodURL 影片H5播放地址
type VodURL struct {
	Model
	VodID           uint        `gorm:"not null;comment:'视频id'"`
	Name            string      `gorm:"comment:'链接名'"`
	URL             string      `gorm:"comment:'链接地址'"`
}

// VodM3u8 影片m3u8地址
type VodM3u8 struct {
	Model
	VodID           uint        `gorm:"not null;comment:'视频id'"`
	Name            string      `gorm:"comment:'链接名'"`
	URL             string      `gorm:"comment:'链接地址'"`
}

// VodDownURL 影片下载地址
type VodDownURL struct {
	Model
	VodID           uint        `gorm:"not null;comment:'视频id'"`
	Name            string      `gorm:"comment:'链接名'"`
	URL             string      `gorm:"comment:'链接地址'"`
}

func autoMigrate() {
	DB.AutoMigrate(&VodKeyword{})
	DB.AutoMigrate(&Vod{})
	DB.AutoMigrate(&VodURL{})
	DB.AutoMigrate(&VodM3u8{})
	DB.AutoMigrate(&VodDownURL{})
}

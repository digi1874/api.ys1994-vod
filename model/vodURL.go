/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-21 21:36:09
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-24 22:54:24
 */

package model

// VodURL vod url
type VodURL struct {
	Base
	ID              uint             `json:"id"`
	VodID           uint             `json:"-"`
	Name            string           `json:"name"`
	URL             string           `json:"url"`
}

// VodM3u8 vod url
type VodM3u8 struct {
	Base
	ID              uint              `json:"id"`
	VodID           uint              `json:"-"`
	Name            string            `json:"name"`
	URL             string            `json:"url"`
}

// VodDownURL vod url
type VodDownURL struct {
	Base
	ID              uint                 `json:"id"`
	VodID           uint                 `json:"-"`
	Name            string               `json:"name"`
	URL             string               `json:"url"`
}

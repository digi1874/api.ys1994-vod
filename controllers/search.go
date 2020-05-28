/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-21 19:32:47
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-26 10:32:55
 */

package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"api.ys1994-vod/database"
	"api.ys1994-vod/model"
	"api.ys1994-vod/utils"
)

// var testListData = `{"code":1,"msg":"\u6570\u636e\u5217\u8868","page":1,"pagecount":1,"limit":"25","total":1,"list":[{"vod_id":47917,"type_id":15,"type_id_1":2,"group_id":0,"vod_name":"\u4f60\u597d\u518d\u89c1\u5988\u5988","vod_sub":"\u54c8\u5570\u63b0\u63b0\uff0c\u6211\u662f\u9b3c\u5988\u5988\/\u518d\u89c1\u5988\u5988\/\u4f60\u597d\u518d\u89c1\uff0c\u5988\u5988\uff01 \/ Hi Bye, Mama! \/ \ud558\uc774\ubc14\uc774,\ub9c8\ub9c8!","vod_en":"nihaozaijianmama","vod_status":1,"vod_letter":"N","vod_color":"","vod_tag":"","vod_class":"\u97e9\u56fd\u5267","vod_pic":"https:\/\/images.cnblogsc.com\/pic\/upload\/vod\/2020-02\/1582398235.jpg","vod_pic_thumb":"","vod_pic_slide":"","vod_actor":"\u91d1\u6cf0\u7199,\u674e\u594e\u70af,\u9ad8\u752b\u6d01,\u7533\u4e1c\u7f8e,\u91d1\u7f8e\u4eac,\u5b89\u5185\u76f8,\u6f58\u5b5d\u8d1e,\u5434\u4e49\u690d,\u88f4\u6d77\u5584","vod_director":"\u67f3\u6d4e\u5143","vod_writer":"","vod_behind":"","vod_blurb":"\u8be5\u5267\u8bb2\u8ff0\u4e16\u754c\u4e0a\u6700\u6696\u5fc3\u7684\u79bb\u522b\u6545\u4e8b\u3002\u63cf\u5199\u5988\u5988\u9b3c\u795e\uff08\u5973\u4e3b\uff09\u4e3a\u4e86\u53d8\u6210\u4eba\u7c7b\u572849\u5929\u91cc\u7684\u771f\u5b9e\u8f6c\u4e16\u8ba1\u5212\uff0c\u548c\u5988\u5988\u9b3c\u795e\u7ecf\u5386\u4e86\u751f\u79bb\u6b7b\u522b\u7684\u75db\u82e6\u5e76\u5c55\u5f00\u65b0\u7684\u4eba\u751f\u7684\u4e08\u592b\uff08\u7537\u4e3b\uff09\uff0c\u4ee5\u53ca\u56de\u5230\u4e24\u4eba\u7684\u5b69\u5b50\uff08\u7ae5\u6f14\uff09\u63cf\u7ed8\u7684\u4e0a\u5929\u548c\u88ab\u7559\u5728\u4eba\u95f4\u7684","vod_remarks":"\u66f4\u65b0\u81f309\u96c6","vod_pubdate":"","vod_total":0,"vod_serial":"9","vod_tv":"","vod_weekday":"","vod_area":"\u97e9\u56fd","vod_lang":"\u97e9\u8bed","vod_year":"2020","vod_version":"","vod_state":"","vod_author":"","vod_jumpurl":"","vod_tpl":"","vod_tpl_play":"","vod_tpl_down":"","vod_isend":0,"vod_lock":0,"vod_level":0,"vod_copyright":0,"vod_points":0,"vod_points_play":0,"vod_points_down":0,"vod_hits":0,"vod_hits_day":0,"vod_hits_week":0,"vod_hits_month":0,"vod_duration":"","vod_up":0,"vod_down":0,"vod_score":"0.0","vod_score_all":0,"vod_score_num":0,"vod_time":"2020-03-23 23:57:03","vod_time_add":1582398364,"vod_time_hits":0,"vod_time_make":0,"vod_trysee":0,"vod_douban_id":0,"vod_douban_score":"0.0","vod_reurl":"","vod_rel_vod":"","vod_rel_art":"","vod_pwd":"","vod_pwd_url":"","vod_pwd_play":"","vod_pwd_play_url":"","vod_pwd_down":"","vod_pwd_down_url":"","vod_content":"\u8be5\u5267\u8bb2\u8ff0\u4e16\u754c\u4e0a\u6700\u6696\u5fc3\u7684\u79bb\u522b\u6545\u4e8b\u3002\u63cf\u5199\u5988\u5988\u9b3c\u795e\uff08\u5973\u4e3b\uff09\u4e3a\u4e86\u53d8\u6210\u4eba\u7c7b\u572849\u5929\u91cc\u7684\u771f\u5b9e\u8f6c\u4e16\u8ba1\u5212\uff0c\u548c\u5988\u5988\u9b3c\u795e\u7ecf\u5386\u4e86\u751f\u79bb\u6b7b\u522b\u7684\u75db\u82e6\u5e76\u5c55\u5f00\u65b0\u7684\u4eba\u751f\u7684\u4e08\u592b\uff08\u7537\u4e3b\uff09\uff0c\u4ee5\u53ca\u56de\u5230\u4e24\u4eba\u7684\u5b69\u5b50\uff08\u7ae5\u6f14\uff09\u63cf\u7ed8\u7684\u4e0a\u5929\u548c\u88ab\u7559\u5728\u4eba\u95f4\u7684\u4eba\u4eec\u7684\u6545\u4e8b\u3002 \u91d1\u6cf0\u7199\u5c06\u5728\u5267\u4e2d\u9970\u6f14\u8f66\u5ba5\u5229\uff0c\u5979\u6027\u683c\u4e50\u5929\u968f\u548c\u70ed\u5fc3\uff0c\u5979\u4e3b\u653b\u73bb\u7483\u5de5\u827a\uff0c\u7ecf\u8425\u7740\u4e00\u4e2a\u5c0f\u4f5c\u574a\uff0c\u5a5a\u540e\u5728\u5750\u6708\u5b50\u7684\u65f6\u5019\u4e5f\u6ca1\u6709\u653e\u5f03\u5de5\u4f5c\u3002\u53bb\u4e16\u7684\u90a3\u5929\u4e5f\u662f\u5728\u53bb\u5de5\u4f5c\u7684\u9014\u4e2d\u9047\u5230\u4e86\u4ea4\u901a\u4e8b\u6545\u3002 \u7531\u300aGo Back\u592b\u5987\u300b\u6743\u6167\u73e0\u7f16\u5267\u6267\u7b14\uff0c\u9884\u8ba1\u63a5\u6863\u300a\u7231\u7684\u8feb\u964d\u300b\u64ad\u51fa\u3002","vod_play_from":"kuyun$$$ckm3u8","vod_play_server":"no$$$no","vod_play_note":"$$$","vod_play_url":"\u7b2c01\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/1d033f2517f2915cbb88cf0a4ceb2611#\u7b2c02\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/22781293bd688d958f3be27e4c26d2c3#\u7b2c03\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/404a9f61135c5e33d002f3fd97350b8f#\u7b2c04\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/cb4e345e4246a6264a5050d32e227b79#\u7b2c05\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/bb9caf96076d335da9c6d368c385dcd3#\u7b2c06\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/ccbaefcc48cd5f8ec9309165ea694eb2#\u7b2c07\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/dba0ff02313bd467ce9d52df8d6c80e6#\u7b2c08\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/b16d68771e66bcf6386df99368154934#\u7b2c09\u96c6$https:\/\/youku.cdn1-okzy.com\/share\/fab6cdec3712c24619019567ac26c060$$$\u7b2c01\u96c6$https:\/\/youku.cdn1-okzy.com\/20200222\/13068_b01af7f1\/index.m3u8#\u7b2c02\u96c6$https:\/\/youku.cdn1-okzy.com\/20200223\/13127_c5841b09\/index.m3u8#\u7b2c03\u96c6$https:\/\/youku.cdn1-okzy.com\/20200229\/13351_bbe7af60\/index.m3u8#\u7b2c04\u96c6$https:\/\/youku.cdn1-okzy.com\/20200306\/13631_911c41bd\/index.m3u8#\u7b2c05\u96c6$https:\/\/youku.cdn1-okzy.com\/20200308\/13701_b3b89896\/index.m3u8#\u7b2c06\u96c6$https:\/\/youku.cdn1-okzy.com\/20200308\/13700_5cb4767c\/index.m3u8#\u7b2c07\u96c6$https:\/\/youku.cdn1-okzy.com\/20200314\/13993_29e02ca3\/index.m3u8#\u7b2c08\u96c6$https:\/\/youku.cdn1-okzy.com\/20200315\/14031_54b31048\/index.m3u8#\u7b2c09\u96c6$https:\/\/youku.cdn1-okzy.com\/20200321\/14260_3cc77fff\/index.m3u8","vod_down_from":"xunlei","vod_down_server":"no","vod_down_note":"","vod_down_url":"\u7b2c01\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200222\/13068_b01af7f1\/\u518d\u89c1\u5988\u598801.mp4#\u7b2c02\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200223\/13127_c5841b09\/\u518d\u89c1\u5988\u598802.mp4#\u7b2c03\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200229\/13351_bbe7af60\/\u518d\u89c1\u5988\u598803.mp4#\u7b2c04\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200306\/13631_911c41bd\/\u518d\u89c1\u5988\u598804.mp4#\u7b2c05\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200308\/13701_b3b89896\/\u518d\u89c1\u5988\u598805.mp4#\u7b2c06\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200308\/13700_5cb4767c\/\u518d\u89c1\u5988\u598806.mp4#\u7b2c07\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200314\/13993_29e02ca3\/\u518d\u89c1\u5988\u598807.mp4#\u7b2c08\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200315\/14031_54b31048\/\u518d\u89c1\u5988\u598808.mp4#\u7b2c09\u96c6$http:\/\/okxzy.xzokzyzy.com\/20200321\/14260_3cc77fff\/\u518d\u89c1\u5988\u598809.mp4","type_name":"\u97e9\u56fd\u5267"}]}`

var okzURL    = "https://api.okzy.tv/api.php/provide/vod/at/json/?ac=detail&"
var httpRe     = regexp.MustCompile(`http(s?)://`)
var scriptRe   = regexp.MustCompile(`<script.+script>`)
var linefeedRe = regexp.MustCompile(`<br ?/?>`)
var hashRe     = regexp.MustCompile(`#`)
var dollarRe   = regexp.MustCompile(`\$`)

type okzVod struct {
	Oid       uint   `json:"vod_id"`
	TypeID    uint   `json:"type_id"`
	TypePID   uint   `json:"type_id_1"`
	Name      string `json:"vod_name"`
	SubName   string `json:"vod_sub"`
	PY        string `json:"vod_en"`
	Pic       string `json:"vod_pic"`
	Actor     string `json:"vod_actor"`
	Director  string `json:"vod_director"`
	Serial    string `json:"vod_remarks"`
	Area      string `json:"vod_area"`
	Lang      string `json:"vod_lang"`
	Year      string `json:"vod_year"`
	Content   string `json:"vod_content"`
	PlayURL   string `json:"vod_play_url"`
	DownURL   string `json:"vod_down_url"`
}

type okzList struct {
	List      []okzVod
	Pagecount int
}

func search(kw string, await bool) bool {
	vKw := model.NewVodKeyword()
	vKw.Filter.Text = kw
	vKw.FirstOrCreate()
	vKw.IncrementNum()

	// 同步中
	if vKw.Filter.SyncLoading == 2 {
		if await {
			awaitSearch(kw)
		}
		return await
	}

	// 同步未超过1小时
	if vKw.Filter.SyncTime > 0 && vKw.Filter.SyncTime + 3600 > uint(time.Now().Unix()) {
		return false
	}

	if await {
		searchStart(*vKw)
	} else {
		go searchStart(*vKw)
	}
	return await
}

func awaitSearch(kw string) {
	loopNum := 0
	Loop:
	vKw := model.NewVodKeyword()
	vKw.Filter.Text = kw
	vKw.FirstOrCreate()
	if vKw.Filter.SyncLoading == 2 && loopNum < 5 {
		time.Sleep(2* time.Second)
		loopNum ++
		goto Loop
	}
}

func searchStart(vKw model.VodKeyword) {
	kw := model.NewVodKeyword()
	kw.Filter.ID = vKw.Filter.ID
	kw.Filter.SyncLoading = 2
	kw.Update()
	defer searchEnd(kw)
	searchOKZ(vKw.Filter.Text, 1)
}

func searchEnd(kw *model.VodKeyword) {
	kw.Filter.SyncLoading = 1
	kw.Filter.SyncTime = uint(time.Now().Unix())
	kw.Update()
}

// searchOKZ 处理搜索OK资源
func searchOKZ (keyword string, page int) bool {
	if keyword == "" {
		return false
	}
	var res okzList
	// err := json.Unmarshal([]byte(testListData), &res)
	url := okzURL + "wd=" + keyword + "&pg=" + strconv.Itoa(page)
	err := utils.GetHTTPJSON(url, &res)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}

	updateData(res.List)

	if res.Pagecount < 10 && res.Pagecount > page {
		return searchOKZ(keyword, page + 1)
	}

	return true
}

func updateData(data []okzVod) {
	for _, value := range data {
		kuYunURL := ""
		ckmURL := ""
		urlGroup := strings.Split(value.PlayURL, "$$$")
		// 提取在线链接；value.URL = "htmlUrls$$$m3u8Urls"
		if strings.Contains(urlGroup[0], "index.m3u8") == false {
			kuYunURL = urlGroup[0]
			if len(urlGroup) > 1 {
				ckmURL = urlGroup[1]
			}
		} else {
			ckmURL = urlGroup[0]
			if len(urlGroup) > 1 {
				kuYunURL = urlGroup[1]
			}
		}

		// 转换一些数据
		value.Content = scriptRe.ReplaceAllString(value.Content, "")
		value.Name = scriptRe.ReplaceAllString(value.Name, "")
		value.Content = linefeedRe.ReplaceAllString(value.Content, "\n")

		// 转json字符串
		if kuYunURL != "" {
			kuYunURL = httpRe.ReplaceAllString(kuYunURL, "")
			kuYunURL = hashRe.ReplaceAllString(kuYunURL, `"},{"name":"`)
			kuYunURL = dollarRe.ReplaceAllString(kuYunURL, `","url":"`)
			kuYunURL = `[{"name":"` + kuYunURL + `"}]`
		}
		if ckmURL != "" {
			ckmURL = httpRe.ReplaceAllString(ckmURL, "")
			ckmURL = hashRe.ReplaceAllString(ckmURL, `"},{"name":"`)
			ckmURL = dollarRe.ReplaceAllString(ckmURL, `","url":"`)
			ckmURL = `[{"name":"` + ckmURL + `"}]`
		}
		if value.DownURL != "" {
			value.DownURL = hashRe.ReplaceAllString(value.DownURL, `"},{"name":"`)
			value.DownURL = dollarRe.ReplaceAllString(value.DownURL, `","url":"`)
			value.DownURL = `[{"name":"` + value.DownURL + `"}]`
		}

		vod := model.NewVod()
		vod.Filter.Oid = value.Oid
		vod.Detail()
		if vod.ID == 0 {
			// 新数据
			err := utils.CopyFromNames(value, &vod.Filter, []string{"TypeID", "TypePID", "Name", "SubName", "PY", "Pic", "Actor", "Director", "Serial", "Area", "Lang", "Year", "Content"})
			if err != nil {
				fmt.Println("CopyFromNames vod.Filter ERROR: ", err)
				return
			}
			if kuYunURL != "" {
				err = json.Unmarshal([]byte(kuYunURL), &vod.Filter.URLs)
				if err != nil {
					fmt.Println("json.Unmarshal vod.Filter.URLs ERROR: ", err)
					return
				}
			}
			if ckmURL != "" {
				err = json.Unmarshal([]byte(ckmURL), &vod.Filter.M3u8s)
				if err != nil {
					fmt.Println("json.Unmarshal vod.Filter.M3u8s ERROR: ", err)
					return
				}
			}
			if value.DownURL != "" {
				err = json.Unmarshal([]byte(value.DownURL), &vod.Filter.DownURLs)
				if err != nil {
					fmt.Println("json.Unmarshal vod.Filter.DownURLs ERROR: ", err)
					return
				}
			}
			vod.Create()
		} else if vod.Serial != value.Serial {
			// 有更新
			updateVod := model.NewVod()
			err := utils.CopyFromNames(value, &updateVod.Filter, []string{"Serial","Pic"})
			if err != nil {
				fmt.Println("CopyFromNames vod.Filter ERROR: ", err)
				return
			}

			if kuYunURL != "" {
				var urls []database.VodURL
				err = json.Unmarshal([]byte(kuYunURL), &urls)
				if err != nil {
					fmt.Println("json.Unmarshal urls ERROR: ", err)
					return
				}
				urlsB, _ := json.Marshal(vod.URLs)
				urlsStr := string(urlsB)
				for _, newValue := range urls {
					if strings.Contains(urlsStr, newValue.URL) == false {
						updateVod.Filter.URLs = append(updateVod.Filter.URLs, newValue)
					}
				}
			}

			if ckmURL != "" {
				var m3u8s []database.VodM3u8
				err = json.Unmarshal([]byte(ckmURL), &m3u8s)
				if err != nil {
					fmt.Println("json.Unmarshal m3u8s ERROR: ", err)
					return
				}
				m3u8sB, _ := json.Marshal(vod.M3u8s)
				m3u8sStr := string(m3u8sB)
				for _, newValue := range m3u8s {
					if strings.Contains(m3u8sStr, newValue.URL) == false {
						updateVod.Filter.M3u8s = append(updateVod.Filter.M3u8s, newValue)
					}
				}
			}

			if value.DownURL != "" {
				var downURLs []database.VodDownURL
				err = json.Unmarshal([]byte(value.DownURL), &downURLs)
				if err != nil {
					fmt.Println("json.Unmarshal downURLs ERROR: ", err)
					return
				}
				downURLsB, _ := json.Marshal(vod.DownURLs)
				downURLsStr := string(downURLsB)
				for _, newValue := range downURLs {
					if strings.Contains(downURLsStr, newValue.URL) == false {
						updateVod.Filter.DownURLs = append(updateVod.Filter.DownURLs, newValue)
					}
				}
			}

			updateVod.Filter.ID = vod.ID
			updateVod.Update()
		}
	}
}

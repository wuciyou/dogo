package main

import (
	"encoding/json"
	"fmt"
	"github.com/wuciyou/dogo"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	get_group_info_ext2 = "http://s.web2.qq.com/api/get_group_info_ext2"
)

type minfo_entity struct {
	City     string
	CounTry  string
	Gender   string
	Nick     string
	Province string
	Uin      int64
}

func getGroupInfoExt2(ctx *dogo.Context) {

	client := http.DefaultClient
	gcode := ctx.Get("gcode")
	vfwebqq := ctx.Get("vfwebqq")
	url := fmt.Sprintf("%s?gcode=%s&vfwebqq=%s&t=%d", get_group_info_ext2, gcode, vfwebqq, (time.Now().UnixNano() / 1000000))
	dogo.Dglog.Debugf("url:%s", url)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	header := request.Header

	header.Add("Cookie", "pgv_pvi=5846345728; pgv_si=s3434502144; ptisp=ctc; RK=g3l+wrQ+sY; ptcz=b3cfd2299cde0cacbf5069a02bf2ed3d644e1c8504b7302cf8c53c333d32d28c; pt2gguin=o3407202930; uin=o3407202930; skey=@AzjNVD9fu; p_uin=o3407202930; p_skey=iplJYRi5miUYCcn65SZoStlqYShVM0*XiUDX6*ovheQ_; pt4_token=OS9Q1s7*7OZ*qmS4VPUSHW1XVApqCXthjGxcivVNecY_")
	header.Add("Host", "s.web2.qq.com")
	header.Add("Referer", "http://s.web2.qq.com/proxy.html?v=20130916001&callback=1&id=1")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36")

	response, err := client.Do(request)

	if err != nil {
		dogo.Dglog.Errorf("请求群信息失败：%v", err)
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		dogo.Dglog.Errorf("获取群数据失败：%v", err)
	}
	var body_json_map map[string]interface{}
	json.Unmarshal(body, &body_json_map)
	dogo.Dglog.Debugf("body_json_map:%+v", body_json_map)
	ctx.W.Write(body)
	ctx.W.Send()
}

func checkUser() func(minfo []minfo_entity) map[string][]minfo_entity {
	var old_minfo map[int64]minfo_entity

	return func(minfo []minfo_entity) map[string][]minfo_entity {
		cur_minfo := make(map[int64]minfo_entity)
		if old_minfo == nil {
			old_minfo = make(map[int64]minfo_entity)
			for _, info := range minfo {
				cur_minfo[info.Uin] = info
			}
			old_minfo = cur_minfo
			return nil
		}
		result := make(map[string][]minfo_entity)
		for uin, info := range old_minfo {
			if _, exist := cur_minfo[uin]; !exist {
				result["quit"] = append(result["quit"], info)

			}
		}
		for uin, info := range cur_minfo {
			if _, exist := old_minfo[uin]; !exist {
				result["new"] = append(result["new"], info)
			}
		}
		return result
	}

}

func loginScan(ctx *dogo.Context) {

}

func main() {
	app := dogo.App()
	app.Route().Get("/login", loginScan)
	app.Route().Get("/get_group_info", getGroupInfoExt2)
	app.Run()
}

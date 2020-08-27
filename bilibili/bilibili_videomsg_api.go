package bilibili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 公有方法
// 获取视频的tag
func GetVideoTag(aid float64) ([]map[string]interface{}, error) {
	// https://api.bilibili.com/x/web-interface/view/detail/tag?aid=754352228
	url := "https://api.bilibili.com/x/web-interface/view/detail/tag?aid=" + strconv.FormatFloat(aid, 'f', -1, 64)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	data, _, err := getDataFromResp(resp, true)

	return data, nil
}

// 通过BVID获取AID和CID
// 同一个视频的bvid和aid是相同的，cid是分P的，没P的cid不同
func GetAidFromBvid(bvId string) (aid float64, cids []float64, bvid string, err error) {
	url := "https://api.bilibili.com/x/web-interface/view?bvid=" + bvId

	//req, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	fmt.Print(err)
	//}
	// 没值赋值，有值不赋值
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	// 有没有值都赋值
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	//client := http.Client{}
	//resp, err := client.Do(req)

	resp, err := http.Get(url)

	if err != nil || resp == nil || resp.StatusCode != 200 {
		return -1, nil, bvId, err
	}

	_, data, err := getDataFromResp(resp, false)
	if err != nil {
		fmt.Print(err)
		return -1, nil, bvid, err
	}

	tmp, err := json.Marshal(data["pages"])
	if err != nil {
		fmt.Print(err)
		return -1, nil, bvid, err
	}

	var pages []map[string]interface{}
	err = json.Unmarshal(tmp, &pages)
	if err != nil {
		fmt.Print(err)
		return -1, nil, bvid, err
	}

	//不用的变量用_表示
	for _, page := range pages {
		cids = append(cids, page["cid"].(float64))
	}

	return data["aid"].(float64), cids, bvid, nil

}

// 私有方法
// {code:0, data:{}/[] }
// 因为resp返回格式大都如此，所以这里对resp解析取出里面的data信息
func getDataFromResp(resp *http.Response, isArray bool) ([]map[string]interface{}, map[string]interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return nil, nil, err
	}

	respMessage := make(map[string]interface{})
	err = json.Unmarshal(body, &respMessage)
	if err != nil {
		fmt.Print(err)
		return nil, nil, err
	}
	dataByte, err := json.Marshal(respMessage["data"])

	if isArray {
		var data []map[string]interface{}
		err = json.Unmarshal(dataByte, &data)
		if err != nil {
			fmt.Print(err)
			return nil, nil, err
		}
		return data, nil, err
	} else {
		data := make(map[string]interface{})
		err = json.Unmarshal(dataByte, &data)
		if err != nil {
			fmt.Print(err)
			return nil, nil, err
		}
		return nil, data, err
	}
}

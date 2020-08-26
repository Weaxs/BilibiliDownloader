package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 通过BVID获取AID和CID
// 同一个视频的bvid和aid是相同的，cid是分P的，没P的cid不同
func getAidFromBvid(bvId string) (aid float64, cids []float64, bvid string) {
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
		return -1, nil, bvId
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, bvid
	}
	respMessage := make(map[string]interface{})
	err = json.Unmarshal(body, &respMessage)
	if err != nil {
		return -1, nil, bvid
	}
	tmp, err := json.Marshal(respMessage["data"])
	data := make(map[string]interface{})
	err = json.Unmarshal(tmp, &data)
	if err != nil {
		return -1, nil, bvid
	}
	tmp, err = json.Marshal(data["pages"])
	var pages []map[string]interface{}
	err = json.Unmarshal(tmp, &pages)
	if err != nil {
		return -1, nil, bvid
	}
	//不用的变量用_表示
	for _, page := range pages {
		cids = append(cids, page["cid"].(float64))
	}

	return data["aid"].(float64), cids, bvid

}

func main() {
	bvid := "BV1ap4y197tR"
	aid, cids, bvid := getAidFromBvid(bvid)
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	fmt.Print("aid=" + strconv.FormatFloat(aid, 'f', -1, 64) + "\ncid= | ")

	for i, cid := range cids {
		fmt.Print(strconv.Itoa(i+1) + "P  " + strconv.FormatFloat(cid, 'f', -1, 64) + " | ")
	}

}

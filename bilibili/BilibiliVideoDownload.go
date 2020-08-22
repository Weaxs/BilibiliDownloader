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
func getAidFromBvid(bvId string) (aid int64, cids []int64, bvid string) {
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

	var cids []int64
	for i, page := range pages {
		fmt.Print("P" + strconv.Itoa(i))
		cids = append(cids, page["cid"].(int64))
	}

	return -1, cids, bvid

}

func main() {
	fmt.Print("")
	bvid := "BV1ap4y197tR"
	aid, cid, bvid := getAidFromBvid(bvid)

	print("aid=" + strconv.FormatInt(aid, 10) + "   cid=" + fmt.Sprint(cid))

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func getAidFromBvid(bvId string) string {
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

	if err != nil || resp == nil {
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if body == nil || err != nil {
		return ""
	}
	respMessage := make(map[string]interface{})
	err = json.Unmarshal(body, &respMessage)
	if err != nil {
		return ""
	}
	body, err = json.Marshal(respMessage["data"])
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ""
	}
	reflect.TypeOf(data["aid"])

	return fmt.Sprint()

}

func main() {
	fmt.Print("")
	aid := "BV1x54y1e7zf"
	if strings.HasPrefix(aid, "BV") {
		getAidFromBvid(aid)
	}

}

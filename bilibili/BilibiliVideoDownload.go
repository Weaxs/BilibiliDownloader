package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getAidFromBvid(bvId string) string {
	url := "https://api.bilibili.com/x/web-interface/view?bvid=" + bvId
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)

	return ""

}

func main() {
	fmt.Print("")
	aid := "BV1x54y1e7zf"
	if strings.HasPrefix(aid, "BV") {
		getAidFromBvid(aid)
	}

}

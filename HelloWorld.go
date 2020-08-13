package main

import (
	"fmt"
	"net/http"
)

func main() {
	var test string = "HelloWorld\n"
	fmt.Print(test)
	fmt.Print("HelloWorld")

}

// 返回error类型
func test() error {

	// 用户输入av号或者视频链接地址
	avNo := "BV1Pi4y1g7Hs"
	start_url := "https://api.bilibili.com/x/web-interface/view?aid=" + avNo
	// map类型
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36"
	delete(headers, "")
	request, err := http.NewRequest(http.MethodGet, start_url, nil)
	if err != nil {
		fmt.Println("new request failed with error: %s", err)
		return err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	client := http.Client{}
	_, err = client.Do(request) // 调用rest接口
	if err != nil {
		return err
	}

	return nil
}

//返回string类型
func testString() string {
	a := "asd"
	switch a {
	case "a":
		return a

	}

	return ""
}

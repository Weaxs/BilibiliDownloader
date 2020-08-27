package bilibili

import (
	"strconv"
	"testing"
)

func Test_GetVideoTag(t *testing.T) {
	if datas, error := GetVideoTag(968064084); error != nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", error) //记录一些你期望记录的信息
		var tag_names []string
		for _, data := range datas {
			tag_names = append(tag_names, data["tag_name"].(string))
		}
		t.Log(tag_names)
	}
}

func Test_GetAidFromBvid(t *testing.T) {
	bvid := "BV1ap4y197tR"
	aid, cids, bvid, error := GetAidFromBvid(bvid)
	if error == nil { //try a unit test on function
		t.Log("one test passed.", error) //记录一些你期望记录的信息
		//fmt.Print("aid=" + strconv.FormatFloat(aid, 'f', -1, 64) + "\ncid= | ")
		t.Log("aid=" + strconv.FormatFloat(aid, 'f', -1, 64) + "\ncid= | ")
		for i, cid := range cids {
			//fmt.Print(strconv.Itoa(i+1) + "P  " + strconv.FormatFloat(cid, 'f', -1, 64) + " | ")
			t.Log(strconv.Itoa(i+1) + "P  " + strconv.FormatFloat(cid, 'f', -1, 64) + " | ")

		}
	} else {
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	}
}

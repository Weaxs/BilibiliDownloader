package main

import (
	"GoRelated/bilibili"
	"fmt"
	"strconv"
)

func main() {
	bvid := "BV1ap4y197tR"
	aid, cids, bvid, _ := bilibili.GetAidFromBvid(bvid)
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

	data, _ := bilibili.GetVideoTag(aid)
	fmt.Print(data)

}

package main

import (
	"auto-youth-study-go/utils"
	"fmt"
)

func main() {
	fmt.Println("开始发送请求......")
	sign := utils.GetSign("xxxxxxx") // mid = "xxxxxxx"
	token := utils.GetToken(sign)
	cid := utils.GetChapterId()
	res := utils.SaveHistory(token, cid)
	if res.Errno == 0 && res.Errmsg == "成功" {
		fmt.Println("保存学习记录成功")
	} else {
		fmt.Println("保存学习记录失败 ", res.Errmsg)
	}
}

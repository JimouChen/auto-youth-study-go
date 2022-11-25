package main

import (
	"auto-youth-study-go/utils"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("请确保你输入了mid值")
		return
	}
	mid := args[len(args)-1]
	fmt.Println("开始发送请求......")
	sign := utils.GetSign(mid) // mid = "xxxxxxx"
	token := utils.GetToken(sign)
	cid := utils.GetChapterId()
	res := utils.SaveHistory(token, cid)
	if res.Errno == 0 && res.Errmsg == "成功" {
		fmt.Println("保存学习记录成功")
	} else {
		fmt.Println("保存学习记录失败", res.Errmsg)
	}
}

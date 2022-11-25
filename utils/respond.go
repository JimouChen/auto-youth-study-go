package utils

import (
	"auto-youth-study-go/params"
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"strings"
)

type SignBody struct {
	YouthLearningUrl string
	Status           int
}

var headers = map[string]string{
	"X-Litemall-Token":          "",
	"X-Litemall-IdentiFication": "young",
	"User-Agent":                "User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c2d) NetType/4G Language/zh_CN",
	"Content-Type":              "application/x-www-form-urlencoded",
}

func GetSign(mid string) string {
	soup.Headers = headers
	urlStr := "https://tuanapi.12355.net/questionnaire/getYouthLearningUrl?mid=" + mid
	resp, err := soup.Get(urlStr)
	if err != nil {
		fmt.Println("get sign failed: ", err.Error())
		panic(err)
	}
	signBody := SignBody{}
	json.Unmarshal([]byte(resp), &signBody)

	return strings.Split(signBody.YouthLearningUrl, "?")[1]
}

func GetToken(sign string) string {
	urlStr := "https://youthstudy.12355.net/apih5/api/user/get"
	soup.Headers = headers
	resp, err := soup.Post(urlStr, "application/x-www-form-urlencoded", sign)
	if err != nil {
		fmt.Println("get token failed: ", err.Error())
		panic(err)
	}
	ts := params.TokenStruct{}
	json.Unmarshal([]byte(resp), &ts)

	return ts.Data.Entity.Token
}

func GetChapterId() string {
	urlStr := "https://youthstudy.12355.net/apih5/api/young/chapter/new"
	soup.Headers = headers
	resp, err := soup.Get(urlStr)
	if err != nil {
		fmt.Println("get chapter id failed: ", err.Error())
		panic(err)
	}
	cidT := params.CidStruct{}
	json.Unmarshal([]byte(resp), &cidT)
	fmt.Println("最新一期是：", cidT.Data.Entity.Name)

	return cidT.Data.Entity.Id
}

func SaveHistory(token string, cid string) (saveRes params.SaveResult) {
	urlStr := "https://youthstudy.12355.net/apih5/api/young/course/chapter/saveHistory"
	headers["X-Litemall-Token"] = token
	soup.Headers = headers
	data := "chapterId=" + cid
	resp, err := soup.Post(urlStr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println("save history failed: ", err.Error())
		panic(err)
	}
	saveRes = params.SaveResult{}
	json.Unmarshal([]byte(resp), &saveRes)

	return
}

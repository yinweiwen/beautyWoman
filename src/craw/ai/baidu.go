package ai

import (
	"craw/craw/utils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type BaiduAi struct {
	token BaiduAiToken
}

func NewBaiduAi() *BaiduAi {
	token, err := BaiduLogin()
	if err != nil {
		panic(err)
	}
	return &BaiduAi{token: token}
}

func (b *BaiduAi) IsAuth() bool {
	return len(b.token.AccessToken) > 0
}

func (b *BaiduAi) Recognize(file string) (ret BaiduAiRecognizeResult, err error) {
	if !b.IsAuth() {
		err = errors.New("not authed")
		return
	}
	var host = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"
	var accessToken = b.token.AccessToken
	var uri *url.URL
	uri, err = url.Parse(host)
	if err != nil {
		return
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()
	filebytes, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	//sendBody.Form.Add("url", "https://baidu-ai.bj.bcebos.com/image-classify/animal.jpeg&baike_num=5")
	sendData := sendBody.Form.Encode()
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(result))
	err = json.Unmarshal(result, &ret)
	return
}

func BaiduLogin() (token BaiduAiToken, err error) {
	var host = "https://aip.baidubce.com/oauth/2.0/token"
	clientId := utils.GIniParser.GetString("ai", "key")
	clientSecret := utils.GIniParser.GetString("ai", "secret")
	fmt.Printf("ai id: %s,secret: %s\n", clientId, clientSecret)
	var param = map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	var uri *url.URL
	uri, err = url.Parse(host)
	if err != nil {
		return
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		return
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(result))
	err = json.Unmarshal(result, &token)
	fmt.Println("get baidu-ai token: ", token.AccessToken)
	return
}

/*
	"refresh_token": "25.b55fe1d287227ca97aab219bb249b8ab.315360000.1798284651.282335-8574074",
	"expires_in": 2592000,
	"scope": "public wise_adapt",
	"session_key": "9mzdDZXu3dENdFZQurfg0Vz8slgSgvvOAUebNFzyzcpQ5EnbxbF+hfG9DQkpUVQdh4p6HbQcAiz5RmuBAja1JJGgIdJI",
	"access_token": "24.6c5e1ff107f0e8bcef8c46d3424a0e78.2592000.1485516651.282335-8574074",
	"session_secret": "dfac94a3489fe9fca7c3221cbf7525ff"
*/
type BaiduAiToken struct {
	RefreshToken  string                 `json:"refresh_token"`
	ExpiresIn     int                    `json:"expires_in"`
	Scope         string                 `json:"scope"`
	AccessToken   string                 `json:"access_token"`
	SessionKey    string                 `json:"session_key"`
	SessionSecret string                 `json:"session_secret"`
	X             map[string]interface{} `json:"-"`
}

/*
log_id	Long	唯一的log id，用于问题定位
result_num	Integer	返回结果数目，及result数组中的元素个数
result	List<Result>	标签结果数组
ResultItem字段数据结构说明
参数名称	参数类型	描述	示例值
keyword	String	图片中的物体或场景名称
score	BigDecimal	置信度，0-1
root	String	识别结果的上层标签，有部分钱币、动漫、烟酒等tag无上层标签
*/
type BaiduAiRecognizeResult struct {
	LogId     int64                        `json:"log_id"`
	ResultNum int                          `json:"result_num"`
	Result    []BaiduAiRecognizeResultItem `json:"result"`
	X         map[string]interface{}       `json:"-"`
}
type BaiduAiRecognizeResultItem struct {
	Keyword string                 `json:"keyword"`
	Score   float64                `json:"score"`
	Root    string                 `json:"root"`
	X       map[string]interface{} `json:"-"`
}

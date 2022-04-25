package dzy

import (
	"craw/craw/sink"
	"craw/craw/utils"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"net"
	"net/http"
	"strings"
	"time"
)

/**
地质云crawer

GET http://171.34.52.5:10026/api/V3.0/Server/GetMonitorPointPartInfo?monitorPointCode=361002010027

Authorization: 3e100cc1-e74f-428e-b66f-d179f7c8a837,b3e22434-3be0-4849-bb25-6d5ea21b6a8e

POST http://171.34.52.5:10026/api/V3.0/Server/GetMonitorPointListSimplePost


请求网址: http://171.34.52.5:10026/Login/CheckLogin
请求方法: POST

__RequestVerificationToken: ZldJ5Xp4dUcwBzKQlxMwZ7N69F9Xi6laLxR0zA5QtXWOjoGFWRiI0lwKq-K5EXjUQuCoV9th-Gkp5F8CUZk9ZsMSMeidAH28lyASHXVp8XM1
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
表单
username: jxfsyj
password: F7A5F37AB905B9DF231484321BD6496F368A2BC42F5D5A0D367BF23AD18C82111D8F1825FDCAA15E
verifycode:
isNoLogin: false

<script>
        function request(d) { for (var c = location.search.slice(1).split("&"), a = 0; a < c.length; a++) { var b = c[a].split("="); if (b[0] == d) if ("undefined" == unescape(b[1])) break; else return unescape(b[1]) } return "" };
        $.rootUrl = '/'.substr(0, '/'.length - 1);
        $.lrToken = $('<input name="__RequestVerificationToken" type="hidden" value="jyyTi_iw0MjFuYu4WQLJs62S0ELRc7bBxhVp6y8_44LWvVOpgDFIbnuR0Dc0x5lTBnOQP00yclTF6KQYr-wy2MNd7a3xVyKJ5Tw-prXIl_E1" />').val();
        //var obj = new WxLogin({
        //    self_redirect: false,
        //    id: "wechatimg",
        //    appid: "wxd75b6651c86f816b",
        //    scope: "snsapi_login",
        //    redirect_uri: "http%3a%2f%2fwww.learun.cn",
        //    state: "",
        //    style: "",
        //    href: ""
        //});


    </script>

password = encryptByDES($.md5(password));
            lrPage.logining("lr_login_btn", true);
            $('#isNoLogin').removeAttr('disabled');
            $.ajax({
                url: $.rootUrl + "/Login/CheckLogin",
                headers: { __RequestVerificationToken: $.lrToken },
                data: { username: username, password: password, verifycode: verifycode, isNoLogin: isNoLogin },
                type: "post",
                dataType: "json",
                success: function (res) {
                    if (res.code == 200) {
                        //密码检查
                        //var IsEnableCheckPwdStrong = $('#IsEnableCheckPwdStrong').val();
                        //if (IsEnableCheckPwdStrong == 1) {
                        //    var passwordCheck = $.trim($password.val());
                        //    var pwModel = checkStrong(passwordCheck);
                        //    if (pwModel < 4) {
                        //        //密码简单，调转修改密码
                        //        tip('当前密码过于简单,请修改,密码必须包含数字,大小写字母和特殊字符');
                        //        lrPage.logining(false);
                        //        $("#lr-login-bypsw").hide();
                        //        $("#lr-register-form").hide();
                        //        $("#txt_find_user").val($("#lr_username").val());
                        //        $("#lr-findPsd-form").show();
                        //        return;
                        //    }
                        //}
                        // DY 修改
                        console.log("tokenMark", res.data);
                        var backUrl = $.cookie('BackURL');
                        if (backUrl == null || backUrl == "" || backUrl == undefined) {
                            backUrl = getQueryStringByName('BackURL');
                        }
                        if (backUrl != null && backUrl != "" && backUrl != undefined) {
                            var connChar = backUrl.indexOf('?') > -1 ? '&' : '?';
                            $.cookie('BackURL', "", { path: "/", expires: 0 });
                            window.location.href = backUrl + connChar + "TokenMark=" + res.data + location.hash;
                        } else {
                            window.location.href = $.rootUrl + '/Home/Index';
                        }
                    }
*/

const (
	domain    = "171.34.52.5"
	loginPage = "http://171.34.52.5:10026/Login/Index?BackURL=http://171.34.52.5:10025/login"
	loginUrl  = "http://171.34.52.5:10026/Login/CheckLogin"
	lstUrl    = "http://171.34.52.5:10026/api/V3.0/Server/GetMonitorPointListSimplePost"
	dataUrl   = "http://171.34.52.5:10026/api/V3.0/Server/GetMonitorPointPartInfo?monitorPointCode=%v"
)

func Craw() {
	curTime := time.Now()

	esUrl := utils.GIniParser.GetString("es", "url")
	username := utils.GIniParser.GetString("dzy", "username")
	password := utils.GIniParser.GetString("dzy", "password")

	c := colly.NewCollector(
		//colly.Async(),
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36"),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second,
			KeepAlive: 300 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	// 初始化值(无用值)
	token := "c0099a3e-3376-45a5-85e7-d506828e3090"
	mark := "139cb977-e713-4c4c-9868-5faa2ed5d2fb"
	__RequestVerificationToken := "YCy9mW0dFZCBeHriVF0ffnULTIsS40bWGmS07zOJ8XXDzuEvx-ZzZRtrxdvbTgTO545MNHJSG_WqEao7OLWCngGApu5aI8JmeKG6EP0Sk6Q1"

	c.SetCookies(domain, []*http.Cookie{
		&http.Cookie{
			Name: "BackURL", Value: "http://171.34.52.5:10025/login",
		},
		&http.Cookie{
			Name: "ASP.NET_SessionId", Value: "50a5lgrlnz2o1hh340btjwzx",
		},
	})

	c.OnRequest(func(r *colly.Request) {
		c.SetCookies(domain, []*http.Cookie{
			&http.Cookie{
				Name: "Learun_ADMS_V7_Mark", Value: mark,
			}, &http.Cookie{
				Name: "Learun_ADMS_V7_Token", Value: token,
			}, &http.Cookie{
				Name: "__RequestVerificationToken", Value: __RequestVerificationToken,
			},
		})
		r.Headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("__RequestVerificationToken", __RequestVerificationToken)
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("Origin", "http://171.34.52.5:10026")
		r.Headers.Set("Referer", "http://171.34.52.5:10026/Login/Index?BackURL=http://171.34.52.5:10025/login")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})

	c.OnHTML("input[name=__RequestVerificationToken]", func(e *colly.HTMLElement) {
		__RequestVerificationToken = e.Attr("value")
	})

	_ = c.Visit(loginPage)
	c.Wait()

	// {"code":200,"info":"登录成功","data":"cdde3ef4-4cc0-4907-b4f6-4dc5196f878c,139cb977-e713-4c4c-9868-5faa2ed5d2fb,jxfsyj"}
	c.OnResponse(func(r *colly.Response) {
		li := LoginInfo{}
		if err := json.Unmarshal(r.Body, &li); err != nil {
			panic(err)
		}
		token, mark = li.getToken()
	})
	//loginContent := "username=jxfsyj&password=&verifycode=&isNoLogin=false"
	c.Post(loginUrl, map[string]string{
		"username":   username,
		"password":   password,
		"verifycode": "",
		"isNoLogin":  "false",
	})
	c.Wait()

	// copy without callback
	c2 := c.Clone()
	c2.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Authorization", token+","+mark)
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})

	c2.OnResponse(func(r *colly.Response) {
		ret := CommRet{}
		json.Unmarshal(r.Body, &ret)

		c3 := c2.Clone()
		c3.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Authorization", token+","+mark)
			r.Headers.Set("Accept", "application/json, text/plain, */*")
			r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		})
		var es *sink.ESClient
		var err error
		if es, err = sink.NewEsClient(esUrl); err != nil {
			panic(err)
		}
		c3.OnResponse(func(r2 *colly.Response) {
			ret := FinalRet{}
			if err := json.Unmarshal(r2.Body, &ret); err == nil {
				ret.ResultData["_code"] = r2.Request.URL.Query().Get("monitorPointCode")
				ret.ResultData["_time"] = curTime
				es.IndexDzy(ret.ResultData)
			}
		})

		for i, datum := range ret.ResultData {
			code := datum["a"]
			url := fmt.Sprintf(dataUrl, code)
			println("[%d] visit [%s]", i, url)
			_ = c3.Visit(url)
			c3.Wait()

			time.Sleep(100 * time.Millisecond)
		}
	})
	content := `{"key":"","queryJson":"{\"keyWords\":\"\",\"MonitorType\":\"\",\"AreaCode\":\"360000\",\"LabelId\":\"\",\"searchDisaster\":\"\",\"WarnLevel\":\"\",\"WeatherLevel\":\"\",\"SYSTEMRESULT\":\"\",\"ISHX\":\"0\",\"ISONLINE\":\"\",\"INDUCEFACTORS\":\"\",\"ModelParam\":\"\",\"DEVICETYPECN\":\"\",\"AndOr\":\"And\",\"XQDJ\":\"\",\"WXRKSTART\":\"\",\"WXRKEND\":\"\",\"WXCCSTART\":\"\",\"WXCCEND\":\"\",\"HAZARDOBJECT\":\"\",\"ISVIRTUAL\":\"0\"}"}`
	_ = c2.PostRaw(lstUrl, []byte(content))
	c2.Wait()

	time.Sleep(10 * time.Minute)
}

type LoginInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data string `json:"data"`
}

type CommRet struct {
	ResultData []map[string]interface{} `json:"resultdata"`
}

type FinalRet struct {
	ResultData map[string]interface{} `json:"resultdata"`
}

/**
,{
    "type":1,
    "resultcode":0,
    "message":"Success",
    "resultdata":{
        "length":"23",
        "width":"54",
        "volume":"0.9936",
        "THICKOPTION":"勘探实测",
        "THICKNUMBER":"8",
        "AREA":null,
        "DITCHLENGTH":null,
        "THREATSPOPULATION":8,
        "THREATTORESIDENTS":2,
        "THREATSASSETS":25,
        "DisasterType":"滑坡",
        "MonitorFre":"",
        "GeoFeauter":null,
        "SetDate":null,
        "DOMURL":null,
        "DOMPAPAM":null,
        "MODELPARAM":null,
        "MODELURL":null,
        "AreaName":null,
        "MonitorPointName":"乐安县金竹乡深坵村王泥坑组滑坡",
        "Location":"江西省抚州市乐安县",
        "Model3DUrl":null,
        "MONITORTYPES":null,
        "LATITUDE":27.01583333,
        "LONGITUDE":115.93305556,
        "ELEVATION":602,
        "MANUFACTURER":"",
        "protectplan":null,
        "contactdeptname":null,
        "contactpeoplename":null,
        "contactphone":null,
        "disasterunitname":null,
        "monitorprovincename":"江西省",
        "monitorcountyname":"乐安县",
        "monitortownname":null,
        "BUILDUNIT":"抚州市自然资源局",
        "MONITORUNIT":"抚州市自然资源局",
        "MONITORTYPE":"滑坡",
        "FIRSTMONITORMAN":null,
        "FIRSTMONITORPHONE":null,
        "MONITORWRANINGMAN":"危斯敏",
        "MONITORWRANINGPHONE":"15979032879",
        "FIRSTPREMAN":null,
        "FIRSTPREMANPHONE":null,
        "SurveyPeople":"曾玉林",
        "SurveyPeoplePhone":"13767626454",
        "NormalCount":4,
        "UnNormalCount":0,
        "WarnLevel":null,
        "IsMonitorxzrecorde":"1",
        "YJXYPEOPLE":"陈正平",
        "YJXYPEOPLEPHONE":"18379496745",
        "FZZRR":"黄爱民",
        "FZZRRDH":"13979469022",
        "JCDTYPE":null,
        "ISHX":"0",
        "HXTYPE":null,
        "DANGERLEVEL":null,
        "CJZRR":null,
        "CJZRRSJ":null,
        "ZJZRR":null,
        "ZJZRRSJ":null,
        "XJZRR":null,
        "XJZRRSJ":null,
        "SCALE":"",
        "MANUFACTURERID":"",
        "CONSTRUCTION":null,
        "MONITORCITYCODE":"361000",
        "PROVINCEMONITORCODE":"361025120078",
        "MONITORCITYNAME":"抚州市",
        "INDUCEFACTORS":"降雨",
        "THREATENEDHOMES":null,
        "HAZARDOBJECT":""
    },
    "otherinfo":null
}
*/

func (l LoginInfo) getToken() (string, string) {
	datas := strings.Split(l.Data, ",")
	return datas[0], datas[1]
}

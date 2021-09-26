package im

import (
	"github.com/parnurzeal/gorequest"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"sports_service/server/app/config"
	"sports_service/server/util"
	"errors"
)

const (
	// 根据运行环境 走配置
	//TX_IM_APP_ID    = 1400570443
	//TX_IM_APP_KEY   = "a6380a7413ed7fdcfac951b9c5fde542c661864c596a4aa7bc1d15c87df5b1f7"
	// 控制台配置的管理员
	//TX_IDENTIFIER   = "bluetrans"
	TX_IM_HOST      = "https://console.tim.qq.com"
)

// 生成签名
func GenSig(userId string, expireTm int) (string, error) {
	sig, err := tencentyun.GenUserSig(config.Global.TencentImAppId, config.Global.TencentImSecret, userId, expireTm)
	if err != nil {
		return "", err
	}

	return sig, nil
}

// 生成请求url
func GenRequestUrl(sig, uri string) string {
	values := url.Values{}
	values.Add("sdkappid", fmt.Sprint(config.Global.TencentImAppId))
	values.Add("identifier", config.Global.TencentImIdentifier)
	values.Add("random", fmt.Sprint(util.GenerateRandnum(100000, 999999)))
	values.Add("usersig", sig)
	values.Add("contenttype", "json")

	return fmt.Sprintf("%s%s?%s", TX_IM_HOST, uri, values.Encode())
}

// post请求
func HttpPostBody(url string, msg map[string]interface{}) ([]byte, error) {
	request := gorequest.New()
	resp, body, errs := request.Post(url).Set("Content-Type", "application/json; charset=utf-8").SendMap(msg).End()
	log.Printf("resp:%+v", resp)
	if errs != nil {
		log.Printf("im_trace: request err:%+v, body:%s", errs, body)
		return nil, errors.New("request error")
	}

	return ioutil.ReadAll(resp.Body)
}

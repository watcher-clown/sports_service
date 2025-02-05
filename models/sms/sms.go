package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sports_service/dao"
	"sports_service/global/app/log"
	"sports_service/global/consts"
	"sports_service/global/rdskey"
	"sports_service/tools/tencentCloud"
	"time"
)

var (
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type SmsModel struct {
}

// 发送短信验证码请求参数
type SendSmsCodeParams struct {
	SendType  string `json:"send_type" binding:"required"`  // 短信类型 1 账户登陆/注册
	MobileNum string `json:"mobile_num" binding:"required"` // 手机号码
}

// 手机验证码登陆 请求参数
type SmsCodeLoginParams struct {
	MobileNum string `binding:"required" json:"mobile_num"` // 手机号码
	Code      string `binding:"required" json:"code"`       // 手机验证码
	Platform  int    `json:"platform"`                      // 平台 0 android 1 iOS 2 web
}

// 实栗
func NewSmsModel() *SmsModel {
	return new(SmsModel)
}

// 获取验证码
func (m *SmsModel) GetSmsCode() string {
	return fmt.Sprintf("%06d", randSource.Intn(999999))
}

// 获取发送验证码的模版
func (m *SmsModel) GetSendMod(sendType string) string {
	switch sendType {
	// 账户登陆/注册 短信模版
	case consts.ACCOUNT_OPT_TYPE:
		return consts.ACCOUNT_MODE
	default:
		log.Log.Errorf("sms_trace: unsupported sendType, sendType:%s", sendType)
	}

	return ""
}

const (
	TEMPLATE_CODE = "SMS_000042" // fpv短信模版code
)

// 已废弃
// 发送短信验证码
//func (m *SmsModel) Send(mobileNum, code string) error {
//	s := &notify.Sms{}
//	s.Mobile = mobileNum
//	s.TemplateParams = m.GetTemplateParams(code)
//	s.TemplateCode = TEMPLATE_CODE
//	s.Time = time.Now().Unix()
//	s.ServiceName = consts.SERVICE_NAME
//	if err := s.Send(); err != nil {
//		return err
//	}
//
//	return nil
//}

// 发送短信验证码
func (m *SmsModel) Send(mobileNum, code string) error {
	client := tencentCloud.New(consts.TX_SMS_SECRET_ID, consts.TX_SMS_SECRET_KEY, consts.TMS_API_DOMAIN)
	rsp, err := client.SendSms(mobileNum, code)
	if err != nil {
		log.Log.Errorf("sms_trace: send sms fail, rsp:%+v, err:%s", rsp, err)
		return err
	}

	if len(rsp.Response.SendStatusSet) > 0 {
		if *rsp.Response.SendStatusSet[0].Code != "Ok" {
			log.Log.Errorf("sms_trace: send sms fail, mobile:%s, rsp:%+v, code:%s", mobileNum, rsp, *rsp.Response.SendStatusSet[0].Code)
			return errors.New("send sms fail")
		}
	}

	return nil
}

type TemplateParams struct {
	Code string `json:"code"`
}

func (m *SmsModel) GetTemplateParams(code string) string {
	params := &TemplateParams{
		Code: code,
	}

	bts, _ := json.Marshal(params)
	return string(bts)
}

// 获取24小时内发送短信的限制数量
func (m *SmsModel) GetSendSmsLimitNum(mobileNum string) (int, error) {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_NUM, time.Now().Format("2006-01-02"), mobileNum)
	rds := dao.NewRedisDao()
	return rds.GetInt(key)
}

// 增加已发短信的数量（24小时内限制十条）
func (m *SmsModel) IncrSendSmsNum(mobileNum string) error {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_NUM, time.Now().Format("2006-01-02"), mobileNum)
	rds := dao.NewRedisDao()
	_, err := rds.INCR(key)
	return err
}

// 记录短信验证码次数的key设置过期
func (m *SmsModel) SetSmsIntervalExpire(mobileNum string) (int, error) {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_NUM, time.Now().Format("2006-01-02"), mobileNum)
	rds := dao.NewRedisDao()
	return rds.EXPIRE(key, rdskey.KEY_EXPIRE_DAY)
}

// 存储验证码 并 设置短信验证码过期时间 5分钟内有效
func (m *SmsModel) SaveSmsCodeByRds(sendType, mobileNum, code string) error {
	key := rdskey.MakeKey(rdskey.SMS_CODE, sendType, mobileNum)
	rds := dao.NewRedisDao()
	return rds.SETEX(key, rdskey.KEY_EXPIRE_MIN*5, code)
}

// 如果redis能获取到 说明验证码未过期
func (m *SmsModel) GetSmsCodeByRds(sendType, mobileNum string) (string, error) {
	key := rdskey.MakeKey(rdskey.SMS_CODE, sendType, mobileNum)
	rds := dao.NewRedisDao()
	return rds.Get(key)
}

// 删除存储验证码的key
func (m *SmsModel) DelSmsCodeKey(sendType, mobileNum string) error {
	key := rdskey.MakeKey(rdskey.SMS_CODE, sendType, mobileNum)
	rds := dao.NewRedisDao()
	_, err := rds.Del(key)
	return err
}

// 是否已过重发验证码的间隔时间
func (m *SmsModel) HasSmsIntervalPass(sendType, mobileNum string) (bool, error) {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_TM, sendType, mobileNum)
	rds := dao.NewRedisDao()
	return rds.EXISTS(key)
}

// 设置重发验证码的间隔时间
func (m *SmsModel) SetSmsIntervalTm(sendType, mobileNum string) error {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_TM, sendType, mobileNum)
	rds := dao.NewRedisDao()
	return rds.SETEX(key, rdskey.KEY_EXPIRE_MIN, 1)
}

// 删除限制重发验证码间隔时间的key
func (m *SmsModel) DelSmsIntervalTmKey(sendType, mobileNum string) error {
	key := rdskey.MakeKey(rdskey.SMS_INTERVAL_TM, sendType, mobileNum)
	rds := dao.NewRedisDao()
	_, err := rds.Del(key)
	return err
}

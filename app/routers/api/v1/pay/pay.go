package pay

import (
	"github.com/gin-gonic/gin"
	wxCli "github.com/go-pay/gopay/wechat"
	"io/ioutil"
	"net/http"
	"net/url"
	"sports_service/server/app/controller/corder"
	"sports_service/server/app/controller/cpay"
	"sports_service/server/global/app/errdef"
	"sports_service/server/global/app/log"
	"sports_service/server/global/consts"
	"sports_service/server/models/morder"
	"sports_service/server/tools/alipay"
	"sports_service/server/tools/wechat"
	"sports_service/server/util"
	"strconv"
	"strings"
	"time"
)

// app发起支付
func AppPay(c *gin.Context) {
	reply := errdef.New(c)
	param := &morder.PayReqParam{}
	if err := c.BindJSON(param); err != nil {
		log.Log.Errorf("pay_trace: invalid param, params:%+v", param)
		reply.Response(http.StatusBadRequest, errdef.INVALID_PARAMS)
		return
	}

	userId, _ := c.Get(consts.USER_ID)
	param.UserId = userId.(string)
	svc := cpay.New(c)
	code, payParam := svc.AppPay(param)
	if code == errdef.SUCCESS {
		reply.Data["pay_param"] = payParam
	}

	reply.Response(http.StatusOK, code)
}

// 支付宝回调通知
func AliPayNotify(c *gin.Context) {
	req := c.Request
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Log.Errorf("aliNotify_trace: err:%s", err.Error())
		c.String(http.StatusBadRequest, "fail")
		return
	}

	log.Log.Debug("aliNotify_trace: info %s", string(body))
	params, err := url.ParseQuery(string(body))
	if err != nil {
		log.Log.Errorf("aliNotify_trace: err:%s, params:%v", err.Error(), params)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	sign := params.Get("sign")
	params.Del("sign")
	params.Del("sign_type")
	query := params.Encode()
	msg, err := url.QueryUnescape(query)
	if err != nil {
		log.Log.Error("aliNotify_trace: QueryUnescape failed: %s", err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	sign, _ = url.PathUnescape(sign)
	log.Log.Debug("aliNotify_trace: msg:%s, sign:%v", msg, sign)

	orderId := params.Get("out_trade_no")
	svc := corder.New(c)
	order, err := svc.GetOrder(orderId)
	if order == nil || err != nil {
		log.Log.Error("aliNotify_trace: order not found, orderId:%s, err:%s", orderId, err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	if order.Status != consts.PAY_TYPE_WAIT {
		log.Log.Error("aliNotify_trace: order already pay, orderId:%s, err:%s", orderId, err)
		c.String(http.StatusOK, "success")
		return
	}

	cli := alipay.NewAliPay(true)
	ok, err := cli.VerifyData(msg, "RSA2", sign)
	if !ok || err != nil {
		log.Log.Errorf("aliNotify_trace: verify data fail, err:%s", err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	amount, err := strconv.ParseFloat(strings.Trim(params.Get("total_amount"), " "), 64)
	if err != nil {
		log.Log.Errorf("aliNotify_trace: parse float fail, err:%s", err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	if int(amount * 100) != order.Amount {
		log.Log.Error("aliNotify_trace: amount not match, orderAmount:%d, amount:%d", order.Amount, amount * 100)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	status := strings.Trim(params.Get("trade_status"), " ")
	payTime, _ := time.Parse("2006-01-02 15:04:05", params.Get("gmt_payment"))
	tradeNo := params.Get("trade_no")
	if err := svc.AliPayNotify(orderId, string(body), status, tradeNo, payTime.Unix(), consts.PAY_NOTIFY); err != nil {
		c.String(http.StatusInternalServerError, "fail")
		return
	}

	c.String(http.StatusOK, "success")
}


type WXPayNotify struct {
	ReturnCode    string `json:"return_code"`
	ReturnMsg     string `json:"return_msg"`
	Appid         string `json:"appid"`
	MchID         string `json:"mch_id"`
	DeviceInfo    string `json:"device_info"`
	NonceStr      string `json:"nonce_str"`
	Sign          string `json:"sign"`
	ResultCode    string `json:"result_code"`
	ErrCode       string `json:"err_code"`
	ErrCodeDes    string `json:"err_code_des"`
	Openid        string `json:"openid"`
	IsSubscribe   string `json:"is_subscribe"`
	TradeType     string `json:"trade_type"`
	BankType      string `json:"bank_type"`
	TotalFee      int64  `json:"total_fee"`
	FeeType       string `json:"fee_type"`
	CashFee       int64  `json:"cash_fee"`
	CashFeeType   string `json:"cash_fee_type"`
	CouponFee     int64  `json:"coupon_fee"`
	CouponCount   int64  `json:"coupon_count"`
	CouponID0     string `json:"coupon_id_0"`
	CouponFee0    int64  `json:"coupon_fee_0"`
	TransactionID string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Attach        string `json:"attach"`
	TimeEnd       string `json:"time_end"`
}

// 微信回调通知
func WechatNotify(c *gin.Context) {
	bm, err := wxCli.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		log.Log.Errorf("wxNotify_trace: parse notify to bodyMap fail, err:%s", err.Error())
		c.String(http.StatusBadRequest, "fail")
		return
	}

	cli := wechat.NewWechatPay(true)
	ok, err := cli.VerifySign(bm)
	if !ok || err != nil {
		log.Log.Error("wxNotify_trace: sign not match, err:%s", err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	body, _ := util.JsonFast.Marshal(bm)
	log.Log.Debug("wxNotify_trace: body:%s, bm:%+v", string(body), bm)

	if hasExist := util.MapExistBySlice(bm, []string{"return_code", "result_code", "out_trade_no", "total_fee",
		"time_end", "transaction_id"}); !hasExist {
		log.Log.Error("wxNotify_trace: map key not exists")
		c.String(http.StatusBadRequest, "fail")
		return
	}

	if bm["return_code"].(string) != "SUCCESS" || bm["result_code"].(string) != "SUCCESS" {
		log.Log.Errorf("wxNotify_trace: trade not success")
		c.String(http.StatusBadRequest, "fail")
		return
	}

	orderId := bm["out_trade_no"].(string)
	svc := corder.New(c)
	order, err := svc.GetOrder(orderId)
	if order == nil || err != nil {
		log.Log.Error("wxNotify_trace: order not found, orderId:%s, err:%s", orderId, err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	if order.Status != consts.PAY_TYPE_WAIT {
		log.Log.Error("wxNotify_trace: order already pay, orderId:%s, err:%s", orderId, err)
		c.String(http.StatusOK, "SUCCESS")
		return
	}

	totalFee := bm["total_fee"].(string)
	fee, err := strconv.Atoi(totalFee)
	if err != nil {
		log.Log.Error("wxNotify_trace: amount not match, orderId:%s, err:%s", orderId, err)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	if fee != order.Amount {
		log.Log.Error("wxNotify_trace: amount not match, orderAmount:%d, amount:%d", order.Amount, fee)
		c.String(http.StatusBadRequest, "fail")
		return
	}

	payTime, _ := time.Parse("20060102150405", bm["time_end"].(string))
	if err := svc.OrderProcess(orderId, string(body), bm["transaction_id"].(string), payTime.Unix(), consts.PAY_NOTIFY); err != nil {
		log.Log.Errorf("wxNotify_trace: order process fail, err:%s", err)
		c.String(http.StatusInternalServerError, "fail")
		return
	}

	c.String(http.StatusOK, "SUCCESS")
}

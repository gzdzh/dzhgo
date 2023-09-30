package service

import (
	"context"
	"encoding/json"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhFinance/config"

	"time"

	cmsCof "github.com/gzdzh/dzhgo/modules/dzhCms/config"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/utility/logger"

	v1 "github.com/gzdzh/dzhgo/modules/dzhFinance/api/v1"
	memberM "github.com/gzdzh/dzhgo/modules/dzhMember/model"

	"github.com/gzdzh/dzhgo/modules/dzhFinance/model"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type DzhFinancePayService struct {
	*dzhCore.Service
}

var (
	client     *wechat.ClientV3
	MchId      string
	APIv3Key   string
	SerialNo   string
	PrivateKey string
	settingMap *gmap.StrAnyMap
	payMap     g.MapStrStr
)

func init() {

	wxpay := config.WxPay(gctx.New())
	MchId = gconv.String(wxpay["SpMchid"])
	APIv3Key = gconv.String(wxpay["APIv3Key"])
	SerialNo = gconv.String(wxpay["SerialNo"])

}

//  NewClientV3 初始化微信客户端
//	mchid：商户ID
//	serialNo：商户证书的证书序列号
//	apiV3Key：APIv3Key，商户平台获取
//	privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容

// native下单
func (s *DzhFinancePayService) V3Native(ctx context.Context, req *v1.V3PartnerNativeReq) (data interface{}, err error) {

	data, err = V3Native(ctx, req.PayModel)

	return
}

// 支付处理
func V3Native(ctx context.Context, payModel *model.PayModel) (data interface{}, err error) {

	setting := cmsCof.GetSetting()
	settingMap = setting.GMap()
	payMap = g.MapStrStr{
		"appid":     gconv.String(settingMap.Get("appid")),     //普通商户模式
		"mchId":     gconv.String(settingMap.Get("mchId")),     //普通商户模式
		"cAPIv3Key": gconv.String(settingMap.Get("cAPIv3Key")), //普通商户模式
		"cSerialNo": gconv.String(settingMap.Get("cSerialNo")), //普通商户模式

		"spMchid":   gconv.String(settingMap.Get("spMchid")),  //服务商模式
		"spAppid":   gconv.String(settingMap.Get("spAppid")),  //服务商模式
		"subMchId":  gconv.String(settingMap.Get("subMchId")), //服务商模式
		"aPIv3Key":  gconv.String(settingMap.Get("aPIv3Key")), //服务商模式
		"serialNo":  gconv.String(settingMap.Get("serialNo")), //服务商模式
		"notifyUrl": gconv.String(settingMap.Get("notifyUrl")),
	}

	payType := gconv.Int(settingMap.Get("payType"))
	g.Log().Debugf(ctx, "[V3Native] payTppe:%d", payType)
	if payType == 1 {
		data, err = V3TransactionNativeUtil(ctx, payModel, payMap)
	} else {
		data, err = V3PartnerNativeUtil(ctx, payModel, payMap)
	}

	return
}

// 普通商户支付下单
func V3TransactionNativeUtil(ctx context.Context, payModel *model.PayModel, payMap g.MapStrStr) (data interface{}, err error) {

	tempFile := "./resource/public/pem/apiclient_key.pem"
	if gfile.Exists(tempFile) {
		PrivateKey = gfile.GetContents(tempFile)
	}

	logger.Debugf(ctx, "tempFile:%v", tempFile)
	// logger.Debugf(ctx, "MchId:%v, SerialNo:%v, APIv3Key:%v, PrivateKey:%v", MchId, SerialNo, APIv3Key, PrivateKey)
	client, err = wechat.NewClientV3(payMap["mchId"], payMap["cSerialNo"], payMap["cAPIv3Key"], PrivateKey)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	// 将价格转换为以分为单位的整数
	priceInCents := int(payModel.Price * 100)

	description := gconv.String(payModel.SkuType) + "套餐:" + gconv.String(payModel.Name)

	bm := make(gopay.BodyMap)
	bm.Set("appid", payMap["appid"]).
		Set("description", description).
		Set("out_trade_no", payModel.TradeNo).
		Set("time_expire", expire).
		Set("notify_url", payMap["notifyUrl"]).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", priceInCents).
				Set("currency", "CNY")
		})
	logger.Debugf(ctx, "下单 bm : %v", bm)

	wxRsp, err := client.V3TransactionNative(ctx, bm)

	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}
	if wxRsp.Code != 0 {
		g.Log().Error(ctx, wxRsp.Error, err)
		err = gerror.New("wxRsp.Code err")
		return
	}

	g.Log().Infof(ctx, "wxRsp: %+v", wxRsp.Response)

	data = wxRsp.Response.CodeUrl

	return
}

// 服务商支付下单
func V3PartnerNativeUtil(ctx context.Context, payModel *model.PayModel, payMap g.MapStrStr) (data interface{}, err error) {

	tempFile := "./resource/public/pem/apiclient_key.pem"
	if gfile.Exists(tempFile) {
		PrivateKey = gfile.GetContents(tempFile)
	}

	logger.Debugf(ctx, "tempFile:%v", tempFile)
	// logger.Debugf(ctx, "MchId:%v, SerialNo:%v, APIv3Key:%v, PrivateKey:%v", MchId, SerialNo, APIv3Key, PrivateKey)
	client, err = wechat.NewClientV3(payMap["spMchid"], payMap["serialNo"], payMap["aPIv3Key"], PrivateKey)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}

	// tradeNo := util.RandomString(32)
	// g.Log().Debug(ctx, "tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	// 将价格转换为以分为单位的整数
	priceInCents := int(payModel.Price * 100)

	description := gconv.String(payModel.SkuType) + "套餐:" + gconv.String(payModel.Name)

	bm := make(gopay.BodyMap)
	bm.Set("sp_appid", payMap["spAppid"]).
		Set("sp_mchid", payMap["spMchid"]).
		Set("sub_mchid", payMap["subMchId"]).
		Set("time_expire", expire).
		Set("notify_url", payMap["notifyUrl"]).
		Set("out_trade_no", payModel.TradeNo).
		Set("description", description).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", priceInCents).
				Set("currency", "CNY")
		})
	logger.Debugf(ctx, "下单 bm : %v", bm)

	wxRsp, err := client.V3PartnerTransactionNative(ctx, bm)

	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}
	if wxRsp.Code != 0 {
		g.Log().Error(ctx, wxRsp.Error, err)
		err = gerror.New("wxRsp.Code err")
		return
	}

	g.Log().Infof(ctx, "wxRsp: %+v", wxRsp.Response)

	data = wxRsp.Response.CodeUrl

	return
}

// 支付回调
func (s *DzhFinancePayService) PayNotice(ctx context.Context, req *v1.NoticeReq) (data interface{}, err error) {

	g.Log().Warningf(ctx, "Notice req :%v", gjson.MustEncodeString(req))

	setting := cmsCof.GetSetting()
	settingMap = setting.GMap()
	payType := gconv.Int(settingMap.Get("payType"))
	g.Log().Debugf(ctx, "[PayNotice] payTppe:%d", payType)
	if payType == 1 {
		data, err = PayNotice(ctx, req)
	} else {
		data, err = PayNoticePartner(ctx, req)
	}

	return
}

// v3服务商支付回调
func PayNoticePartner(ctx context.Context, req *v1.NoticeReq) (data interface{}, err error) {

	setting := cmsCof.GetSetting()
	settingMap = setting.GMap()
	APIv3Key = gconv.String(settingMap.Get("aPIv3Key"))

	result, err := wechat.V3DecryptPartnerNotifyCipherText(req.Resource.Ciphertext, req.Resource.Nonce, req.Resource.AssociatedData, APIv3Key)
	if err != nil {
		bytes, _ := json.Marshal(req)
		logger.Errorf(ctx, "[V3DecryptPartnerNotifyCipherText](%s) decrypt cipher text error(%s)", string(bytes), err)
		return
	}

	g.Log().Warningf(ctx, "[V3DecryptPartnerNotifyCipherText] result %v", gjson.MustEncodeString(result))

	data, err = payOrderUpdate(ctx, result.OutTradeNo)

	if err != nil {
		g.Log().Error(ctx, "payOrderUpdate error:"+err.Error(), err)
		err = gerror.New("订单状态更新失败")
		return
	}
	return
}

// v3普通商回调
func PayNotice(ctx context.Context, req *v1.NoticeReq) (data interface{}, err error) {

	setting := cmsCof.GetSetting()
	settingMap = setting.GMap()
	APIv3Key = gconv.String(settingMap.Get("cAPIv3Key"))

	result, err := wechat.V3DecryptNotifyCipherText(req.Resource.Ciphertext, req.Resource.Nonce, req.Resource.AssociatedData, APIv3Key)
	if err != nil {
		bytes, _ := json.Marshal(req)
		logger.Errorf(ctx, "[V3DecryptNotifyCipherText] (%s) decrypt cipher text error(%s)", string(bytes), err)
		return
	}

	g.Log().Warningf(ctx, "[V3DecryptNotifyCipherText] result %v", gjson.MustEncodeString(result))

	data, err = payOrderUpdate(ctx, result.OutTradeNo)

	if err != nil {
		g.Log().Error(ctx, "payOrderUpdate error:"+err.Error(), err)
		err = gerror.New("订单状态更新失败")
		return
	}
	return
}

// 支付回调工具函数
func payOrderUpdate(ctx context.Context, outTradeNo string) (data interface{}, err error) {

	tradeNo := outTradeNo

	orderInfo, _ := dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).One()
	if orderInfo.IsEmpty() {
		g.Log().Warning(ctx, "order isEmpty")
		err = gerror.New("没有此订单")
		return
	}

	if gconv.Bool(orderInfo["payStatus"]) {
		g.Log().Warning(ctx, "The order has been paid")
		err = gerror.New("订单已经支付过了")
		return
	}

	userId := gconv.Int64(orderInfo["userId"])
	memberInfo, _ := dzhCore.DBM(memberM.NewMemberUserAttr()).Where("userId", userId).One()

	if !memberInfo.IsEmpty() {
		// 已存在购买过

		// gpt套餐
		if gconv.String(orderInfo["skuType"]) == "GPT" {

			var gptEndDate string
			if memberInfo["gptEndDate"] != nil {
				// 之前购买过
				t1 := gtime.New(gconv.String(memberInfo["gptEndDate"]))
				t2 := gtime.New(time.Now())

				if t2.Before(t1) {
					// 未到期的叠加
					g.Log().Debug(ctx, "未到期的叠加")
					gptEndDate = t1.AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")

				} else {
					// 已到期的重新记录到期时间
					g.Log().Debug(ctx, "已到期的或者第一次购买的重新记录到期时间")
					gptEndDate = gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")
				}
			} else {
				// 第一次购买
				gptEndDate = gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")
			}

			updateMap := g.Map{
				"gptEndDate": gptEndDate,
				"gptSkuName": orderInfo["name"],
				"gptPayTime": gtime.Now(),
			}
			g.Log().Debug(ctx, "updateMap", updateMap)
			_, err = dzhCore.DBM(memberM.NewMemberUserAttr()).Where("userId", userId).Data(updateMap).Update()
			if err != nil {
				g.Log().Error(ctx, "pay notice order update GPT error:"+err.Error(), err)
				err = gerror.New("订单更新失败")
				return
			}

			// 写入支付时间
			dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).Data("gptPayTime", gtime.Now()).Update()
		}

		// MJ套餐
		if gconv.String(orderInfo["skuType"]) == "MJ" {

			// 计算有效期
			var mjEndDate string
			if memberInfo["mjEndDate"] != nil {
				// 之前购买过
				t1 := gtime.New(gconv.String(memberInfo["mjEndDate"]))
				t2 := gtime.New(time.Now())

				if t2.Before(t1) {
					// 未到期的叠加
					g.Log().Debug(ctx, "未到期的叠加")
					mjEndDate = t1.AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")

				} else {
					// 已到期的重新记录到期时间
					g.Log().Debug(ctx, "已到期的或者第一次购买的重新记录到期时间")
					mjEndDate = gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")
				}
			} else {
				// 第一次购买
				mjEndDate = gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")
			}

			// 计算张数
			mjCount := gconv.Int64(orderInfo["mjCount"])
			// 剩余的次数叠加
			if gconv.Int(memberInfo["mjCount"]) > 0 {
				mjCount = mjCount + gconv.Int64(memberInfo["mjCount"])
			}

			updateMap := g.Map{
				"mjCount":   mjCount,
				"mjSkuName": orderInfo["name"],
				"mjEndDate": mjEndDate,
				"mjPayTime": gtime.Now(),
			}
			_, err = dzhCore.DBM(memberM.NewMemberUserAttr()).Where("userId", userId).Data(updateMap).Update()
			if err != nil {
				g.Log().Error(ctx, "pay notice order update MJ error:"+err.Error(), err)
				err = gerror.New("订单更新失败")
				return
			}
			// 写入支付时间
			dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).Data("mjPayTime", gtime.Now()).Update()
		}

	} else {
		// 未曾购买过的

		// GPT套餐
		if gconv.String(orderInfo["skuType"]) == "GPT" {
			// 第一次购买
			gptEndDate := gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")
			updateMap := g.Map{
				"userId":     userId,
				"gptEndDate": gptEndDate,
				"gptSkuName": orderInfo["name"],
				"gptPayTime": gtime.Now(),
			}
			_, err = dzhCore.DBM(memberM.NewMemberUserAttr()).Where("userId", userId).Data(updateMap).Insert()
			if err != nil {
				g.Log().Error(ctx, "pay notice order insert error:"+err.Error(), err)
				err = gerror.New("订单写入失败")
				return
			}
			// 写入支付时间
			dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).Data("gptPayTime", gtime.Now()).Update()
		}

		// MJ套餐
		if gconv.String(orderInfo["skuType"]) == "MJ" {

			// 第一次购买
			mjEndDate := gtime.New(time.Now()).AddDate(0, gconv.Int(orderInfo["validDate"]), 0).Format("Y-m-d")

			mjCount := gconv.Int64(orderInfo["mjCount"])
			updateMap := g.Map{
				"userId":    userId,
				"mjSkuName": orderInfo["name"],
				"mjCount":   mjCount,
				"mjEndDate": mjEndDate,
				"mjPayTime": gtime.Now(),
			}
			_, err = dzhCore.DBM(memberM.NewMemberUserAttr()).Where("userId", userId).Data(updateMap).Insert()
			if err != nil {
				g.Log().Error(ctx, "pay notice order insert error:"+err.Error(), err)
				err = gerror.New("订单写入失败")
				return
			}
			// 写入支付时间
			dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).Data("mjPayTime", gtime.Now()).Update()

		}

	}
	_, err = dzhCore.DBM(model.NewDzhFinanceOrder()).Where("tradeNo", tradeNo).Data("payStatus", 1).Update()
	if err != nil {

		return
	}

	return
}

func PayNoticeT(ctx context.Context, outTradeNo string) (data interface{}, err error) {

	data, err = payOrderUpdate(ctx, outTradeNo)

	return
}

func NewDzhFinancePayService() *DzhFinancePayService {
	return &DzhFinancePayService{
		&dzhCore.Service{},
	}
}

package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
	comS "github.com/gzdzh/dzhgo/modules/dzhCommon/service"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/utility/logger"
	v1 "github.com/gzdzh/dzhgo/modules/dzhFinance/api/v1"
	dzhFinanceModel "github.com/gzdzh/dzhgo/modules/dzhFinance/model"

	// "github.com/gzdzh/dzhgo/modules/dzhFinance/service"
	dzhMembereModel "github.com/gzdzh/dzhgo/modules/dzhMember/model"

	// dzhPayModel "github.com/gzdzh/dzhgo/modules/dzhPay/model"
	// dzhPayService "github.com/gzdzh/dzhgo/modules/dzhPay/service"

	"context"

	"github.com/go-pay/gopay/pkg/util"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type DzhFinanceOrderService struct {
	*dzhCore.Service
}

// ServiceAdd 新增
func (s *DzhFinanceOrderService) ServiceAdd(ctx context.Context, req *dzhCore.AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)

	rmap := r.GetMap()
	// 非空键
	if s.NotNullKey != nil {
		for k, v := range s.NotNullKey {
			if rmap[k] == nil {
				return nil, gerror.New(v)
			}
		}
	}
	// 唯一键
	if s.UniqueKey != nil {
		for k, v := range s.UniqueKey {
			if rmap[k] != nil {
				m := dzhCore.DBM(s.Model)
				count, err := m.Where(k, rmap[k]).Count()
				if err != nil {
					return nil, err
				}
				if count > 0 {
					err = gerror.New(v)
					return nil, err
				}
			}
		}
	}
	if s.InsertParam != nil {
		insertParams := s.InsertParam(ctx)
		if len(insertParams) > 0 {
			for k, v := range insertParams {
				rmap[k] = v
			}
		}
	}
	m := dzhCore.DBM(s.Model)
	lastInsertId, err := m.Data(rmap).InsertAndGetId()
	if err != nil {
		return
	}

	payModel := &dzhFinanceModel.PayModel{
		SkuType: gconv.String(rmap["skuType"]),
		TradeNo: gconv.String(rmap["tradeNo"]),
		Price:   gconv.Float64(rmap["price"]),
		Name:    gconv.String(rmap["name"]),
	}

	logger.Debugf(ctx, "payModel:%v", gjson.MustEncodeString(payModel))

	res, err := V3Native(ctx, payModel)

	if err != nil {
		logger.Error(ctx, "ServiceAdd V3PartnerNativeUtil err", err.Error())
		return
	}

	data = g.Map{
		"id":       lastInsertId,
		"tradeNo":  rmap["tradeNo"],
		"code_url": res,
	}

	return
}

// 前置处理
func (s *DzhFinanceOrderService) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {
	member := comS.GetMember(ctx)
	if method == "Add" {
		sku := gconv.String(param["sku"])
		goodInfo, err := dzhCore.DBM(model.NewCmsNews()).Where("sku", sku).One()
		if err != nil {
			return err
		}
		if goodInfo.IsEmpty() {
			logger.Debugf(ctx, "goodInfo IsEmpty")
			err = gerror.New("goodInfo IsEmpty")
			return err
		}

		var (
			trade   string
			skuType string
		)
		if goodInfo["typeId"].Int() == 1 {
			trade = "gpt_no_"
			skuType = "GPT"

		}
		if goodInfo["typeId"].Int() == 2 {
			trade = "mj_no_"
			skuType = "MJ"
		}
		tradeNo := trade + util.RandomString(7)
		g.RequestFromCtx(ctx).SetParam("userId", member.UserId)
		g.RequestFromCtx(ctx).SetParam("sku", sku)
		g.RequestFromCtx(ctx).SetParam("tradeNo", tradeNo)
		g.RequestFromCtx(ctx).SetParam("name", goodInfo["name"])
		g.RequestFromCtx(ctx).SetParam("price", goodInfo["price"])
		g.RequestFromCtx(ctx).SetParam("validDate", goodInfo["validDate"])
		g.RequestFromCtx(ctx).SetParam("skuType", skuType)
		g.RequestFromCtx(ctx).SetParam("mjCount", goodInfo["mjCount"])

	}
	return
}

// 后置处理
func (s *DzhFinanceOrderService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {

	return
}

// 查询订单支付状态
func (s *DzhFinanceOrderService) OrderPayStatus(ctx g.Ctx, req *v1.OrderPayStatusReq) (data interface{}, err error) {
	member := comS.GetMember(ctx)
	whereMap := g.Map{
		"tradeNo": req.TradeNo,
		"userId":  member.UserId,
	}
	orderInfo, err := dzhCore.DBM(s.Model).Where(whereMap).One()
	if err != nil {
		g.Log().Error(ctx, "finance_order orderStatus find err:"+err.Error(), err)
		err = gerror.New("查询错误")
		return
	}

	if orderInfo.IsEmpty() {
		g.Log().Error(ctx, "finance_order orderStatus find empty")
		err = gerror.New("无订单数据")
		return
	}

	if gconv.Bool(orderInfo["payStatus"]) {
		data = "已支付"
	}

	return
}

func NewDzhFinanceOrderService() *DzhFinanceOrderService {
	return &DzhFinanceOrderService{
		&dzhCore.Service{
			Model:       dzhFinanceModel.NewDzhFinanceOrder(),
			ListQueryOp: &dzhCore.QueryOp{},
			PageQueryOp: &dzhCore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{"username", "gptSkuName", "mjSkuName"},
				AddOrderby:   map[string]string{"`dzh_dzhfinance_order`.`createTime`": "DESC"},
				Select:       "dzh_dzhfinance_order.*,b.username,b.nickname,c.gptSkuName,c.gptEndDate,c.mjSkuName,c.mjEndDate,c.mjCount",
				Join: []*dzhCore.JoinOp{
					{
						Model:     dzhMembereModel.NewMemberUser(),
						Alias:     "b",
						Type:      "LeftJoin",
						Condition: "`dzh_dzhfinance_order`.`userId` = `b`.`id`",
					},
					{
						Model:     dzhMembereModel.NewMemberUserAttr(),
						Alias:     "c",
						Type:      "LeftJoin",
						Condition: "`c`.`userId` = `b`.`id`",
					},
				},
			},
		},
	}
}

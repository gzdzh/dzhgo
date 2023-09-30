package model

import (
	"time"

	"github.com/gzdzh/dzhgo/dzhCore"
)

const TableNameDzhFinanceOrder = "dzh_dzhfinance_order"

// DzhFinanceOrder mapped from table <dzhfinance_order>
type DzhFinanceOrder struct {
	*dzhCore.Model
	Name       string    `gorm:"column:name;not null;comment:名称" json:"name"`
	Img        string    `gorm:"column:img;comment:图片" json:"img"`
	Status     *int32    `gorm:"column:status;comment:状态;default:1" json:"status"`
	Remark     string    `gorm:"column:remark;comment:备注" json:"remark"`
	TradeNo    string    `gorm:"column:tradeNo;comment:订单编号" json:"tradeNo"`
	Price      string    `gorm:"column:price;comment:价格" json:"price"`
	Sku        string    `gorm:"column:sku;not null;comment:sku;index" json:"sku"`
	UserId     int64     `gorm:"column:userId;not null;comment:会员;index" json:"userId"`
	ValidDate  *int      `gorm:"column:validDate;comment:使用期限" json:"validDate"`
	SkuType    *string   `gorm:"column:skuType;not null;comment:产品类型;index" json:"skuType"`
	MjCount    int64     `gorm:"column:mjCount;not null;comment:MJ使用数量;default:0" json:"mjCount"`
	PayStatus  *int32    `gorm:"column:payStatus;comment:支付状态;default:0" json:"payStatus"`
	GptPayTime time.Time `gorm:"column:gptPayTime;index;comment:GPT支付时间" json:"gptPayTime"`
	MJPayTime  time.Time `gorm:"column:mjPayTime;index,;comment:MJ支付时间" json:"mjPayTime"`
}

// TableName DzhFinanceOrder's table name
func (*DzhFinanceOrder) TableName() string {
	return TableNameDzhFinanceOrder
}

// GroupName DzhFinanceOrder's table group
func (*DzhFinanceOrder) GroupName() string {
	return "default"
}

// NewDzhFinanceOrder create a new DzhFinanceOrder
func NewDzhFinanceOrder() *DzhFinanceOrder {
	return &DzhFinanceOrder{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&DzhFinanceOrder{})
}

package model

import (
	"time"

	"github.com/gzdzh/dzhgo/dzhCore"
)

const TableNameMemberUserAttr = "dzh_member_user_attr"

// MemberUserAttr mapped from table <member_user_attr>
type MemberUserAttr struct {
	*dzhCore.Model
	UserId     *string   `gorm:"column:userId;not null;comment:会员" json:"userId"`
	GptEndDate *string   `gorm:"column:gptEndDate;comment:GPT到期时间" json:"gptEndDate"`
	GptPayTime time.Time `gorm:"column:gptPayTime;comment:GPT购买时间" json:"gptPayTime"`

	MjEndDate  *string   `gorm:"column:mjEndDate;comment:MJ到期时间" json:"mjEndDate"`
	GptSkuName *string   `gorm:"column:gptSkuName;comment:gpt套餐" json:"gptSkuName"`
	MjSkuName  *string   `gorm:"column:mjSkuName;comment:mj套餐" json:"mjSkuName"`
	MjCount    int64     `gorm:"column:mjCount;not null;comment:MJ使用数量;default:0" json:"mjCount"`
	MjPayTime  time.Time `gorm:"column:mjPayTime;comment:MJ购买时间" json:"mjPayTime"`
}

// TableName MemberUserAttr's table name
func (*MemberUserAttr) TableName() string {
	return TableNameMemberUserAttr
}

// GroupName MemberUserAttr's table group
func (*MemberUserAttr) GroupName() string {
	return "default"
}

// NewMemberUserAttr create a new MemberUserAttr
func NewMemberUserAttr() *MemberUserAttr {
	return &MemberUserAttr{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&MemberUserAttr{})
}

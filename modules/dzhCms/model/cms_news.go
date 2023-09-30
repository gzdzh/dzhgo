package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsNews = "dzh_cms_news"

// CmsNews mapped from table <cms_news>
type CmsNews struct {
	*dzhCore.Model
	Title    string  `gorm:"column:title;not null;comment:标题" json:"title"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Abstract *string `gorm:"column:abstract;comment:摘要" json:"abstract"`
	Remark   *string `gorm:"column:remark;comment:内容" json:"remark"`
	Click    *int64  `gorm:"column:click;comment:点击;type:bigint;default:0" json:"click"`
	Status   string  `gorm:"column:status;comment:状态;type:int" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`
}

// TableName CmsNews's table name
func (*CmsNews) TableName() string {
	return TableNameCmsNews
}

// GroupName CmsNews's table group
func (*CmsNews) GroupName() string {
	return "default"
}

// NewCmsNews create a new CmsNews
func NewCmsNews() *CmsNews {
	return &CmsNews{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsNews{})
}

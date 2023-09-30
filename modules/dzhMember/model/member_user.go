package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameMemberUser = "dzh_member_user"

// MemberUser mapped from table <member_user>
type MemberUser struct {
	*dzhCore.Model
	AvatarUrl   string  `gorm:"column:avatarUrl;comment:头像;type:varchar(200)" json:"avatarUrl"`
	Username    string  `gorm:"column:username;not null;comment:会员账号;type:varchar(50);index" json:"username"`
	Password    string  `gorm:"column:password;not null;comment:会员密码;type:varchar(50)" json:"password"`
	PasswordV   *int32  `gorm:"column:passwordV;not null;type:int;default:1" json:"passwordV"` // 密码版本, 作用是改完密码，让原来的token失效
	NickName    string  `gorm:"column:nickName;comment:会员昵称;type:varchar(50);index" json:"nickName"`
	LevelName   string  `gorm:"column:levelName;comment:等级;type:varchar(50)" json:"levelName"`
	Sex         *int32  `gorm:"column:sex;comment:性别;type:int;default:1" json:"sex"`
	QQ          *string `gorm:"column:qq;comment:qq;type:varchar(255);index" json:"qq"`
	Mobile      *string `gorm:"column:mobile;comment:手机号;type:varchar(50);index" json:"mobile"`
	Wx          *string `gorm:"column:wx;comment:微信号;type:varchar(50);index" json:"wx"`
	WxImg       *string `gorm:"column:wxImg;comment:微信二维码;type:varchar(255)" json:"wxImg"`
	Email       *string `gorm:"column:email;comment:email;type:varchar(50);index" json:"email"`
	Role        string  `gorm:"column:role;comment:家庭角色;" json:"role"`
	Remark      *string `gorm:"column:remark;comment:备注;type:varchar(255)" json:"remark"`
	Openid      string  `gorm:"column:openid;comment:openid;" json:"openid"`
	Unionid     string  `gorm:"column:unionid;comment:unionid;" json:"unionid"`
	SessionKey  string  `gorm:"column:session_key;comment:session_key;" json:"session_key"`
	Status      *int32  `gorm:"column:status;not null;type:int;default:1" json:"status"` // 状态 0:禁用 1：启用
	Description *string `gorm:"column:description;comment:描述;type:varchar(100)" json:"description"`
}

// TableName MemberUser's table name
func (*MemberUser) TableName() string {
	return TableNameMemberUser
}

// GroupName MemberUser's table group
func (*MemberUser) GroupName() string {
	return "default"
}

// NewMemberUser create a new MemberUser
func NewMemberUser() *MemberUser {
	return &MemberUser{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&MemberUser{})
}

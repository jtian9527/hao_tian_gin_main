package dao

import (
	"gorm.io/gorm"
	"haotian_main/utils"
)

type MemberInfo struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	SellerId int64  `json:"seller_id"`
	BrandUid int64  `json:"brand_uid"`
	BindUid  string `json:"bind_uid"`
	Point    int64  `json:"point"`
	Tier     int64  `json:"tier"`
	TierName string `json:"tier_name"`
	Spending int64  `json:"spending" `
	Area     string `json:"area"`
	IsDelete bool   `json:"is_delete"`
	Ctime    int64  `json:"ctime"`
	Mtime    int64  `json:"mtime"`
	TraceId  string `json:"trace_id"`
}

func (MemberInfo) TableName() string {
	return "demo_member_info"
}

type MemberInfoDao struct {
	gormCluster *gorm.DB
}

var memberInfoDao = &MemberInfoDao{gormCluster: utils.DBCluster}

func GetMemberInfoDao() *MemberInfoDao {
	return memberInfoDao
}

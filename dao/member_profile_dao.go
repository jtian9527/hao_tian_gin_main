package dao

import (
	"gorm.io/gorm"
	"haotian_main/utils"
)

type MemberProfile struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	SellerId  int64  `json:"seller_id"`
	BuyerUid  int64  `json:"buyer_uid"`
	BrandUid  int64  `json:"brand_uid"`
	BindUid   string `json:"bind_uid"`
	BuyerName string `json:"buyer_name"`
	BindTime  int64  `json:"bind_time"`
	Phone     int64  `json:"phone"`
	Birthday  int64  `json:"birthday" `
	Email     string `json:"email"`
	Area      string `json:"area"`
	Extra     string `json:"extra" `
	JoinTime  int64  `json:"join_time"`
	IsDelete  bool   `json:"is_delete"`
	Ctime     int64  `json:"ctime"`
	Mtime     int64  `json:"mtime"`
	TraceId   string `json:"trace_id"`
}

func (MemberProfile) TableName() string {
	return "demo_member_profile"
}

type MemberProfileDao struct {
	gormCluster *gorm.DB
}

var memberProfileDao = &MemberProfileDao{gormCluster: utils.DBCluster}

func GetMemberProfileDao() *MemberProfileDao {
	return memberProfileDao
}

func (p *MemberProfileDao) GetUserProfileInfobySellerIDandPhone(sellerId int64, phone int64) (memberProfileDao []MemberProfile) {
	p.gormCluster.First(&memberProfileDao, "seller_id = ? AND phone = ? AND is_delete=?", sellerId, phone, 0)
	return
}

func (p *MemberProfileDao) GetUserProfileInfobySellerIDandBindUid(sellerId int64, bindUid string) (memberProfileDao []MemberProfile) {
	p.gormCluster.First(&memberProfileDao, "seller_id = ? AND bind_uid = ? AND is_delete=?", sellerId, bindUid, 0)
	return
}

func (p *MemberProfileDao) GetUserProfileInfobyTraceId(traceId string) (memberProfileDao []MemberProfile) {
	p.gormCluster.First(&memberProfileDao, "seller_id = ? AND is_delete=?", traceId, 0)
	return
}
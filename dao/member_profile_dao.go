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

//func (p *MemberProfileDao) UpdateUserInfoProfile(params *param2.BindMemberReq, tierName string, bindUid string) {
//	var dbInfo *gorm.DB
//	tx := dbInfo.Begin()
//	tx.Updates(&MemberProfile{}).Where("seller_id = ? AND phone = ? AND is_delete=?", params.SellerId, params.Phone, 0).
//		Updates(map[string]interface{}{
//			"seller_id": params.SellerId,
//			"buyer_id":  params.BuyerId,
//			"brand_uid": bindUid,
//			"buyer_name":params.BuyerName,
//			"bind_time":time.Now().Unix(),
//			"phone":params.Phone,
//			"birthday":params.Birthday,
//			"email":params.Email,
//			"area":params.Area,
//			"extra":params.Extra,
//			"join_time":params.JoinTime,
//			"is_delete":0,
//			"mtime":time.Now().Unix(),
//			"trace_id":params.TraceId,
//		})
//
//	if dbInfo.Error != nil {
//		return
//	}
//
//	return
//}

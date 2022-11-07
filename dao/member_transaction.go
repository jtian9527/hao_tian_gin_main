package dao

import (
	"gorm.io/gorm"
	param2 "haotian_main/param"
	"haotian_main/utils"
	"time"
)

type MemberTransactionModel interface {
	UpdateUserInfoProfile(params *param2.BindMemberReq, tierName string, bindUid string, brandUid int64) (err error)
	CreateUserInfoProfile(params *param2.BindMemberReq, tierName string, bindUid string, brandUid int64) (err error)
	QuitUserInfoProfile(params *param2.UnbindMemberReq)(err error)
}

type MemberTransactionImpl struct {

}

func (m *MemberTransactionImpl)UpdateUserInfoProfile(params *param2.BindMemberReq, tierName string, bindUid string, brandUid int64) (err error){
	var db *gorm.DB
	tx := db.Begin()
	tx.Updates(&MemberProfile{}).Where("seller_id = ? AND phone = ? AND is_delete=?", params.SellerId, params.Phone, 0).
		Updates(map[string]interface{}{
			"seller_id":  params.SellerId,
			"buyer_id":   params.BuyerUid,
			"brand_uid":  brandUid,
			"bind_id":    bindUid,
			"buyer_name": params.BuyerName,
			"bind_time":  time.Now().Unix(),
			"phone":      params.Phone,
			"birthday":   params.Birthday,
			"email":      params.Email,
			"area":       params.Area,
			"extra":      params.Extra,
			"join_time":  params.JoinTime,
			"is_delete":  false,
			"mtime":      time.Now().Unix(),
			"trace_id":   params.TraceId,
		})
	tx.Updates(&MemberInfo{}).Where("seller_id = ? AND brand_uid = ? AND is_delete=?", params.SellerId, brandUid, 0).
		Updates(map[string]interface{}{
			"seller_id": params.SellerId,
			"brand_uid": brandUid,
			"bind_id":   bindUid,
			"point":     0,
			"tier":      params.Tier,
			"tier_name": tierName,
			"Spending":  params.Spending,
			"area":      params.Area,
			"is_delete": false,
			"mtime":     time.Now().Unix(),
			"trace_id":  params.TraceId,
		})
	tx.Commit()
	return
}

func (m *MemberTransactionImpl)CreateUserInfoProfile(params *param2.BindMemberReq, tierName string, bindUid string, brandUid int64) (err error) {
	var db *gorm.DB
	db = utils.DBCluster
	tx := db.Begin()
	err = tx.Create(&MemberProfile{
		SellerId:  params.SellerId,
		BuyerUid:  params.BuyerUid,
		BrandUid:  brandUid,
		BindUid:   bindUid,
		BuyerName: params.BuyerName,
		BindTime:  time.Now().Unix(),
		Phone:     params.Phone,
		Birthday:  params.Birthday,
		Email:     params.Email,
		Area:      params.Area,
		Extra:     params.Extra,
		JoinTime:  params.JoinTime,
		IsDelete:  false,
		Ctime:     time.Now().Unix(),
		Mtime:     time.Now().Unix(),
		TraceId:   params.TraceId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&MemberInfo{
		SellerId: params.SellerId,
		BrandUid: brandUid,
		BindUid:  bindUid,
		Point:    0,
		Tier:     params.Tier,
		TierName: tierName,
		Spending: params.Spending,
		Area:     params.Area,
		IsDelete: false,
		Ctime:    time.Now().Unix(),
		Mtime:    time.Now().Unix(),
		TraceId:  params.TraceId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *MemberTransactionImpl) QuitUserInfoProfile(params *param2.UnbindMemberReq) (err error) {
	var db *gorm.DB
	db = utils.DBCluster
	tx := db.Begin()
	err = tx.Model(&MemberProfile{}).Where("seller_id = ? AND bind_uid = ? AND is_delete=?", params.SellerId, params.BindUid, 0).Updates(MemberProfile{
		BindUid:  "",
		BuyerUid: 0,
		Mtime:    time.Now().Unix(),
		TraceId:  params.TraceId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Updates(&MemberInfo{}).Where("seller_id = ? AND bind_uid = ? AND is_delete=?", params.SellerId, params.BindUid, 0).Updates(MemberInfo{
		BindUid: "",
		Mtime:   time.Now().Unix(),
		TraceId: params.TraceId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Updates(&PointDetail{}).Where("seller_id = ? AND bind_uid = ? AND is_delete=?", params.SellerId, params.BindUid, 0).Updates(PointDetail{
		IsDelete: true,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

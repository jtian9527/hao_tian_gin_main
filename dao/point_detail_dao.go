package dao

import (
	"gorm.io/gorm"
	"haotian_main/utils"
)

type PointDetail struct {
	ID             int64  `json:"id" gorm:"primaryKey"`
	SellerId       int64  `json:"seller_id"`
	BrandUid       int64  `json:"brand_uid"`
	BindUid        string `json:"bind_uid"`
	ChangePoint    int64  `json:"change_point"`
	OldPoint       int64  `json:"old_point"`
	NewPoint       int64  `json:"new_point"`
	TraceId        string `json:"trace_id"`
	TraceType      string `json:"trace_type" `
	Extra          string `json:"extra" `
	IsSelfOperated bool   `json:"is_self_operated"`
	Area           string `json:"area"`
	IsDelete       bool   `json:"is_delete"`
	Ctime          int64  `json:"ctime"`
}

func (PointDetail) TableName() string {
	return "demo_point_detail"
}

type PointDetailDao struct {
	gormCluster *gorm.DB
}

var pointDetailDao = &PointDetailDao{gormCluster: utils.DBCluster}

func GetPointDetailDao() *PointDetailDao {
	return pointDetailDao
}

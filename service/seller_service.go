package service

import (
	"haotian_main/config"
	param2 "haotian_main/param"
	"haotian_main/constant"
	"haotian_main/dao"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type SellerService struct{}

var sellerService = &SellerService{}

func GetSellerService() *SellerService {
	return sellerService
}

func (r *SellerService) BindMember(ctx *gin.Context, params *param2.BindMemberReq) *param2.BindMemberResp {
	BindMemberResp := &param2.BindMemberResp{}
	//读取配置
	conf, err := config.ParseConfig("./config/config.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	tierName := conf.TierRule[params.Tier-1].TierName

	userProfile := dao.GetMemberProfileDao().GetUserProfileInfobySellerIDandPhone(params.SellerId,params.Phone)
	if len(userProfile) == 1 {
		if userProfile[0].BindUid == ""{
			bindUid := CreateBindUid(params.Phone)
			//这里做好用事务处理
			dao.UpdateUserInfoProfile(params, tierName, bindUid, userProfile[0].BrandUid)
		}else{
			BindMemberResp.Code = constant.HadBind
		}
	}else{
		bindUid := CreateBindUid(params.Phone)
		dao.CreateUserInfoProfile(params, tierName, bindUid, userProfile[0].BrandUid)
	}
	return BindMemberResp
}

func CreateBindUid(phone int64) (bindUid string){
	firstData:=strconv.FormatInt(phone,10)
	secondData:=strconv.FormatInt(time.Now().Unix(),10)
	bindUid = firstData + secondData
	return bindUid
}
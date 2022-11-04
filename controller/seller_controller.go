package controller

import (
	"github.com/gin-gonic/gin"
	"haotian_main/constant"
	param2 "haotian_main/param"
	"haotian_main/service"
)


type SellerController struct {
}

func (r *SellerController) BindMember(ctx *gin.Context) {
	var params *param2.BindMemberReq
	Resp := &param2.BindMemberResp{}
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		Resp.Code = constant.CheckJsonFail
		r.BindMemberResponse(ctx, Resp)
		return
	}
	sellerService := service.GetSellerService()
	Resp = sellerService.BindMember(ctx, params)
	r.BindMemberResponse(ctx, Resp)
}

func (r *SellerController) BindMemberResponse(c *gin.Context, Resp *param2.BindMemberResp) {
	if Resp.Msg == "" {
		Resp.Msg = GetMsgFromCode(Resp.Code)
	}
	c.JSON(200, Resp)
}
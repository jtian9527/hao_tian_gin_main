package param

type BindMemberReq struct {
	SellerId  int64  `json:"seller_id" binding:"required"`
	BuyerUid  int64  `json:"buyer_uid" binding:"required"`
	TraceId   string `json:"trace_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Area      string `json:"area" binding:"required"`
	Phone     int64  `json:"phone" binding:"required"`
	BuyerName string `json:"buyer_name" binding:"required"`
	Birthday  int64  `json:"birthday" binding:"required"`
	Tier      int64  `json:"tier" binding:"required"`
	JoinTime  int64  `json:"join_time" binding:"required"`
	Spending  int64  `json:"spending" binding:"required"`
	Extra     string `json:"extra"`
}
type BindMemberResp struct {
	BindUid string `json:"bind_uid"`
	Code    uint32 `json:"code"`
	Msg     string `json:"msg"`
}

type UnBindMemberReq struct {
	SellerId int64  `json:"seller_id" binding:"required"`
	TraceId  string `json:"trace_id" binding:"required"`
	Area     string `json:"area" binding:"required"`
	BindUid  string `json:"bind_uid" binding:"required"`
}
type UnBindMemberResp struct {
	Data *UnBindData `json:"data"`
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
}
type UnBindData struct {
	Point int64 `json:"point"`
	Tier  int64 `json:"tier"`
}

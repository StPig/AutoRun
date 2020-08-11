package common

// User ...
type User struct {
	LineID             string
	McDailyAccount     string
	McDailyPassword    string
	McDailyAccessToken string
	IPassCardNo        string
}

type memberInfo struct {
	AccessToken string `json:"access_token"`
}

type couponObjectInfo struct {
	Title      string `json:"title"`
	ExpireTime string `json:"redeem_end_datetime"`
}

type stickerObjectInfo struct {
	Title      string `json:"title"`
	ExpireTime string `json:"expire_datetime"`
}

type coupon struct {
	ObjectInfo couponObjectInfo `json:"object_info"`
}

type sticker struct {
	ObjectInfo stickerObjectInfo `json:"object_info"`
}

type loginResults struct {
	MemberInfo memberInfo `json:"member_info"`
}

type lotteryResults struct {
	Coupon coupon `json:"coupon"`
}

type lotteryListResults struct {
	Coupons []coupon `json:"coupons"`
}

type stickerListResults struct {
	Stickers []sticker `json:"stickers"`
}

// LoginResp ...
type LoginResp struct {
	Rc      string       `json:"rc"`
	Rm      string       `json:"rm"`
	Results loginResults `jason:"results"`
}

// LotteryResp ...
type LotteryResp struct {
	Rc      int            `json:"rc"`
	Rm      string         `json:"rm"`
	Results lotteryResults `jason:"results"`
}

// LotteryListResp ...
type LotteryListResp struct {
	Rc      int                `json:"rc"`
	Rm      string             `json:"rm"`
	Results lotteryListResults `jason:"results"`
}

// StickerListResp ...
type StickerListResp struct {
	Rc      int                `json:"rc"`
	Rm      string             `json:"rm"`
	Results stickerListResults `jason:"results"`
}

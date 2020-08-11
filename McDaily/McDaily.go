package mcdaily

import (
	"McDailyAutoRun/common"
	"McDailyAutoRun/config"
	lineapi "McDailyAutoRun/lineAPI"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func getToken(account string, password string) string {
	fmt.Println("start to get token")

	if account == "" || password == "" {
		fmt.Println("account or password is empty")
		return ""
	}

	client := resty.New()

	var (
		deviceTime  = time.Now().Format("2006/01/02 15:04:05")
		callTime    = time.Now().Format("20060102150405")
		paramString = account + password
		orderNo     = config.DeviceUUID + callTime
		maskMd5     = md5.Sum([]byte(fmt.Sprintf("Mc%s%s%s%s%s%s%s%sDonalds",
			orderNo,
			config.Platform,
			config.OsVersion,
			config.ModelID,
			config.DeviceUUID,
			deviceTime,
			config.AppVersion,
			paramString)))
	)

	params := fmt.Sprintf(`
	{"account":"%s",
	"password":"%s",
	"OrderNo":"%s",
	"mask":"%x",
	"source_info":{
		"app_version":"%s",
		"device_time":"%s",
		"device_uuid":"%s",
		"model_id":"%s",
		"os_version":"%s",
		"Platform":"%s"}
	}`,
		account,
		password,
		orderNo,
		maskMd5,
		config.AppVersion,
		deviceTime,
		config.DeviceUUID,
		config.ModelID,
		config.OsVersion,
		config.Platform)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&common.LoginResp{}).
		Post(config.McDailyLoginURL)

	if err != nil {
		fmt.Println("fail to login to McDaily", err.Error())
		return ""
	} else if resp.StatusCode() != 200 {
		fmt.Printf("fail to login to McDaily, status code: %d\n", resp.StatusCode())
		return ""
	}

	r := resp.Result().(*common.LoginResp)

	if r.Rc != "1" {
		fmt.Printf("fail to login to McDaily: %s\n", r.Rm)
		return ""
	}

	return r.Results.MemberInfo.AccessToken
}

// GetLottery ...
func GetLottery(user common.User) {
	fmt.Println("start to get lottery")

	if user.McDailyAccessToken == "" {
		user.McDailyAccessToken = getToken(user.McDailyAccount, user.McDailyPassword)
	}

	client := resty.New()

	deviceTime := time.Now().Format("2006/01/02 15:04:05")

	params := fmt.Sprintf(`
	{"access_token":"%s",
	"source_info":{
		"app_version":"%s",
		"device_time":"%s",
		"device_uuid":"%s",
		"model_id":"%s",
		"os_version":"%s",
		"Platform":"%s"}
	}`,
		user.McDailyAccessToken,
		config.AppVersion,
		deviceTime,
		config.DeviceUUID,
		config.ModelID,
		config.OsVersion,
		config.Platform)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&common.LotteryResp{}).
		Post(config.McDailyGetLottry)

	if err != nil {
		fmt.Println("fail to get lottery", err.Error())
		return
	} else if resp.StatusCode() != 200 {
		fmt.Printf("fail to get lottery, status code: %d\n", resp.StatusCode())
		return
	}

	r := resp.Result().(*common.LotteryResp)

	if r.Rc != 1 {
		fmt.Printf("fail to get lottery: %s\n", r.Rm)
		return
	}

	if r.Rm == "今日已領過，明日驚喜等著你！" {
		fmt.Println(r.Rm)
	} else {
		msg := fmt.Sprintf("恭喜獲得: %s, 有效期限: %s\n", r.Results.Coupon.ObjectInfo.Title, r.Results.Coupon.ObjectInfo.ExpireTime)
		fmt.Printf(msg)
		lineapi.PushMessageToUser(user.LineID, msg)
	}
}

// GetLotteryList ...
func GetLotteryList(user common.User) common.LotteryListResp {
	var lotteryList common.LotteryListResp
	fmt.Println("start to get lottery list")

	if user.McDailyAccessToken == "" {
		user.McDailyAccessToken = getToken(user.McDailyAccount, user.McDailyPassword)
	}

	client := resty.New()

	deviceTime := time.Now().Format("2006/01/02 15:04:05")

	params := fmt.Sprintf(`
	{"access_token":"%s",
	"source_info":{
		"app_version":"%s",
		"device_time":"%s",
		"device_uuid":"%s",
		"model_id":"%s",
		"os_version":"%s",
		"Platform":"%s"}
	}`,
		user.McDailyAccessToken,
		config.AppVersion,
		deviceTime,
		config.DeviceUUID,
		config.ModelID,
		config.OsVersion,
		config.Platform)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&lotteryList).
		Post(config.McDailyGetLottryList)

	if err != nil {
		fmt.Println("fail to get lottery list", err.Error())
		return lotteryList
	} else if resp.StatusCode() != 200 {
		fmt.Printf("fail to get lottery list, status code: %d\n", resp.StatusCode())
		return lotteryList
	} else if lotteryList.Rc != 1 {
		fmt.Printf("fail to get lottery list: %s\n", lotteryList.Rm)
		return lotteryList
	}

	return lotteryList
}

// GetStickerList ...
func GetStickerList(user common.User) common.StickerListResp {
	var stickerList common.StickerListResp
	fmt.Println("start to get sticker list")

	if user.McDailyAccessToken == "" {
		user.McDailyAccessToken = getToken(user.McDailyAccount, user.McDailyPassword)
	}

	client := resty.New()

	deviceTime := time.Now().Format("2006/01/02 15:04:05")

	params := fmt.Sprintf(`
	{"access_token":"%s",
	"source_info":{
		"app_version":"%s",
		"device_time":"%s",
		"device_uuid":"%s",
		"model_id":"%s",
		"os_version":"%s",
		"Platform":"%s"}
	}`,
		user.McDailyAccessToken,
		config.AppVersion,
		deviceTime,
		config.DeviceUUID,
		config.ModelID,
		config.OsVersion,
		config.Platform)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&stickerList).
		Post(config.McDailyGetStickerList)

	if err != nil {
		fmt.Println("fail to get sticker list", err.Error())
		return stickerList
	} else if resp.StatusCode() != 200 {
		fmt.Printf("fail to get sticker list, status code: %d\n", resp.StatusCode())
		return stickerList
	} else if stickerList.Rc != 1 {
		fmt.Printf("fail to get sticker list: %s\n", stickerList.Rm)
		return stickerList
	}

	return stickerList
}

package mpAuth

import (
	"flag"
	"fmt"
	"lemon-wxmp-sdk/core"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
)

var (
	appId  = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	secret = flag.String("WXMP_SECRET", os.Getenv("WXMP_SECRET"), "WXMP_SECRET")
)

func Test_JoinOpenIdParam(t *testing.T) {
	reqDto := &ReqDto{
		AppId: *appId,
		Scope: core.SNSAPI_USERINFO,
		State: "1",

		RedirectUrl: "https://gateway.p2shop.cn/qywxaccount/ping",
		PageUrl:     "",
	}
	fmt.Println(GetUrlForAccessToken(reqDto))
}

/*
test steps:
1.get url:https://xxx by GetUrlForAccessToken
2.get code, by accessing this url:https://xxx  by wechat debug tools
*/
func Test_GetAccessTokenAndOpenId(t *testing.T) {
	code := "021BObOo0XJcys1B2YOo0tuVNo0BObOK"
	respDto, err := GetAccessTokenAndOpenId(code, *appId, *secret)
	test.Ok(t, err)
	fmt.Printf("%+v", respDto)
}

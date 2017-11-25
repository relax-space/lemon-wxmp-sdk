package mpAuth

import (
	"fmt"
	"lemon-wxmp-sdk/core"

	"github.com/relax-space/go-kit/httpreq"
)

type ReqDto struct {
	AppId string `json:"app_id"` //required
	Scope string `json:"scope"`  //option
	State string `json:"state"`  //state

	//Secret      string `json:"secret"`
	RedirectUrl string `json:"redirect_url"`
	PageUrl     string `json:"page_url"`
}

type RespDto struct {
	core.RespErrorDto
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

func GetUrlForAccessToken(dto *ReqDto) (reqUrl string) {
	if len(dto.Scope) == 0 {
		dto.Scope = core.SNSAPI_BASE
	}
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v?reurl=%v&response_type=code&scope=%v&state=%v#wechat_redirect",
		dto.AppId,
		dto.RedirectUrl,
		dto.PageUrl,
		dto.Scope,
		dto.State,
	)
}

func GetAccessTokenAndOpenId(code, appId, secret string) (respDto *RespDto, err error) {
	reqUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code",
		appId,
		secret,
		code,
	)
	respDto = &RespDto{}
	_, err = httpreq.POST("", reqUrl, nil, respDto)
	if err != nil {
		return
	}
	if respDto.ErrCode != 0 {
		err = fmt.Errorf("%v:%v||%v", core.MESSAGE_WECHAT, respDto.ErrCode, respDto.ErrMsg)
		return
	}
	return
}

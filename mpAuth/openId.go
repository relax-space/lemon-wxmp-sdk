package mpAuth

import (
	"fmt"
	"net/url"

	"github.com/relax-space/lemon-wxmp-sdk/core"

	"github.com/relax-space/go-kit/httpreq"
)

type ReqDto struct {
	AppId string `json:"app_id" query:"app_id"` //required
	Scope string `json:"scope" query:"app_id"`  //option
	State string `json:"state" query:"app_id"`  //option

	//Secret      string `json:"secret"`
	RedirectUrl string `json:"redirect_url" query:"app_id"`
	PageUrl     string `json:"page_url" query:"app_id"` //option
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
	if len(dto.State) == 0 {
		dto.State = "state"
	}
	if len(dto.PageUrl) == 0 {
		return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v&response_type=code&scope=%v&state=%v#wechat_redirect",
			dto.AppId,
			url.QueryEscape(dto.RedirectUrl),
			dto.Scope,
			dto.State,
		)
	}
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v?reurl=%v&response_type=code&scope=%v&state=%v#wechat_redirect",
		dto.AppId,
		url.QueryEscape(dto.RedirectUrl),
		url.QueryEscape(dto.PageUrl),
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

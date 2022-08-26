// Code generated by goctl. DO NOT EDIT.
package types

type ExpandReq struct {
	Snowflake string `path:"snowflake"`
}

type ExpandResp struct {
	RedirectUrl string `json:"redirectUrl"`
}

type ShortenReq struct {
	RedirectUrl string `json:"redirectUrl"`
	Secret      string `json:"secret,optional"`
	ExpiresIn   int64  `json:"expiresIn,optional"`
}

type ShortenResp struct {
	Id string `json:"id"`
}

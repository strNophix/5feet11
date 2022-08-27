// Code generated by goctl. DO NOT EDIT.
package types

type ExpandReq struct {
	ID string `path:"id"`
}

type ExpandResp struct {
	RedirectUrl string `json:"redirectUrl"`
}

type ShortenReq struct {
	RedirectUrl string `json:"redirectUrl"`
	ExpiresIn   int64  `json:"expiresIn,optional"`
}

type ShortenResp struct {
	ID string `json:"id"`
}

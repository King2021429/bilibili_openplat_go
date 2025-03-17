package model

const (
	AcceptHeader              = "Accept"
	ContentTypeHeader         = "Content-Type"
	AuthorizationHeader       = "Authorization"
	JsonType                  = "application/json"
	HmacSha256                = "HMAC-SHA256"
	BiliTimestampHeader       = "x-bili-timestamp"
	BiliSignatureMethodHeader = "x-bili-signature-method"
	BiliSignatureNonceHeader  = "x-bili-signature-nonce"
	BiliAccessKeyIdHeader     = "x-bili-accesskeyid"
	BiliSignVersionHeader     = "x-bili-signature-version"
	BiliContentMD5Header      = "x-bili-content-md5"
	AccessToken               = "access-token"

	BiliVersion   = "1.0"
	BiliVersionV2 = "2.0"

	MethodGet  = "GET"
	MethodPost = "POST"
)

type CommonHeader struct {
	ContentType       string
	ContentAcceptType string
	Timestamp         string
	SignatureMethod   string
	SignatureVersion  string
	Authorization     string
	Nonce             string
	AccessKeyId       string
	ContentMD5        string
	AccessToken       string
}

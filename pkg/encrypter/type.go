package encrypter

type DataWithExpire struct {
	Data       string `json:"data"`
	ExpireTime int64  `json:"expire_time"`
}

package model

type Url struct {
	LongUrl  string `json:"longurl"`
	ShortUrl string `json:"shorturl"`
	Key      string `json:"key"`
}

func (u Url) GetLongUrl() string {
	return u.LongUrl
}

func (u Url) GetShortUrl() string {
	return u.ShortUrl
}

func (u Url) GetKey() string {
	return u.Key
}

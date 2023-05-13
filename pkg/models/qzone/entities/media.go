package qzone

type Media struct {
	Type string `json:"type"` // "video" or "image"
	Url  string `json:"url"`
}

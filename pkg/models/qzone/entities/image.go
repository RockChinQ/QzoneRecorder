package qzone

type Image struct {
	ID        int64 // 数据库中的id
	EmotionID int64 // 说说id(数据库)
	Url       string
}

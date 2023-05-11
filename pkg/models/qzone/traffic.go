package qzone

type Traffic struct {
	ID              int64 // 数据库中的id
	EmotionID       int64 // 说说id(数据库)
	Likes           int   // 点赞数
	Visits          int   // 浏览量
	Forwards        int   // 转发数
	UpdateTimeStamp int64 // 更新时间戳
}

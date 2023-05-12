package qzone

type Comment struct {
	ID        int64 // 数据库中的id
	EmotionID int64 // 说说id(数据库)
	PersonQQ  int64 // 评论人QQ
	Text      string
	TimeStamp int64 // 评论时间戳
	ParentID  int64 // 父评论id
}

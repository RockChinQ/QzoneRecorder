package qzone

type Emotion struct {
	ID        int64  // 数据库中的id
	Eid       string // 说说id(腾讯系统)
	PersonID  int64  // 说说所属人的id
	Text      string // 说说文字内容
	TimeStamp int64  // 说说发表时间戳
}

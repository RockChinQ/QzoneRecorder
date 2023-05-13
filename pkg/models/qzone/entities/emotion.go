package qzone

type Emotion struct {
	Eid       string  `json:"eid"`
	QQ        string  `json:"qq"`
	Nickname  string  `json:"nickname"`
	Text      string  `json:"text"`
	Medias    []Media `json:"medias"`
	TimeStamp string  `json:"time_stamp"`
}

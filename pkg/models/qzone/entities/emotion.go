package qzone

type Emotion struct {
	Eid           string    `json:"eid"`
	UserCard      UserCard  `json:"user_card"`
	Text          string    `json:"text"`
	Medias        []Media   `json:"medias"`
	Comments      []Comment `json:"comments"`
	TimeStamp     int       `json:"time_stamp"`
	Traffic       Traffic   `json:"traffic"`
	DetailPageUrl string    `json:"detail_page_url"`
}

package qzone

type Comment struct {
	UserCard  UserCard `json:"user_card"`
	TimeStamp string   `json:"time_stamp"`
	Text      string   `json:"text"`
}

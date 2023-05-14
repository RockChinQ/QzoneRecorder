package qzone

type Traffic struct {
	Likers        []UserCard `json:"likers"`
	ForwardAmount int        `json:"forward_amount"`
	VisitAmount   int        `json:"visit_amount"`
}

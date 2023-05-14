package qzone

type Traffic struct {
	Likers        []UserCard `json:"likers"`
	LikeAmount    int        `json:"like_amount"`
	ForwardAmount int        `json:"forward_amount"`
	VisitAmount   int        `json:"visit_amount"`
	CommentAmount int        `json:"comment_amount"`
}

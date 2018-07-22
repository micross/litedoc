package models

import (
	"time"
)

type Vote struct {
	Id            int       `json:"id"`
	CommentId     int       `json:"comment_id"`
	CommentUserId int       `json:"comment_user_id"`
	VoteUserId    int       `json:"vote_user_id"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

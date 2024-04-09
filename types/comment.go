package types

import (
	"time"

	"github.com/gofrs/uuid"
)

type Comment struct {
	ID           uuid.UUID `json:"uuid"`
	User         User      `json:"user"`
	Post_ID      uuid.UUID `json:"post_id"`
	CreationDate time.Time `json:"creationDate"`
	Content      string    `json:"content"`
	Likes        int64     `json:"likes"`
}

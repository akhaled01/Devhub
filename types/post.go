package types

import (
	"image"
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID           uuid.UUID   `json:"uuid"`
	User         User        `json:"user"`
	Image        image.Image `json:"image"`
	Likes        int         `json:"likes"`
	Comments     []Comment   `json:"comments"`
	Category     Category    `json:"category"`
	CreationDate time.Time   `json:"creationDate"`
}

func CreatePost(){}
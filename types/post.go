package types

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID                 uuid.UUID   `json:"id"`
	User               PartialUser `json:"user"`
	Image_Path         string      `json:"image"`
	Likes              int64       `json:"likes"`
	Liked              bool        `json:"liked"`
	Comments           []Comment   `json:"comments"`
	Category           []string    `json:"category"`
	CreationDate       time.Time   `json:"creationDate"`
	Content            string      `json:"content"`
	Number_of_comments int         `json:"number_of_comments"`
}

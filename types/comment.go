package types

import "github.com/gofrs/uuid"

type Comment struct {
	ID    uuid.UUID `json:"uuid"`
	User  User      `json:"user"`
	Post  Post      `json:"post"`
	Likes int       `json:"likes"`
}

package types

import "github.com/gofrs/uuid"

// to be decoded from a signup request
type SignupRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Gender    string `json:"gender"`
	Avatar    string `json:"image"`
}

type LoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}

type PostCreationRequest struct {
	Session_id        string // use to extract user who created the post
	Post_text         string `json:"post_text"`
	Post_image_base64 string `json:"post_image"`
	Post_category     int    `json:"post_category"`
}

type CommentCreationRequest struct {
	Session_id   string
	Post_id      uuid.UUID `json:"post_id"`
	Comment_text string    `json:"comment_text"`
}

package types

import "errors"

var (
	ErrPrepare           = errors.New("error preparing sql statement")
	ErrAppendPost        = errors.New("error appending post")
	ErrPostNotFound      = errors.New("post not found")
	ErrScan              = errors.New("error scanning rows")
	ErrExec              = errors.New("error executing statemnt")
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrGetLikes          = errors.New("error getting likes")
	ErrUUID              = errors.New("error dealing with uuid")
	ErrGetCommentDetails = errors.New("error getting comment details")
	ErrCats              = errors.New("error getting full category details")
	ErrImage             = errors.New("error dealing with images correctly")
	ErrSessionNotFound   = errors.New("session not found")
)

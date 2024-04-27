package comment

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"
)

const SaveCommentInDBQuery = `INSERT INTO comments (comm_id, post_id, user_id, comment_date, comment) VALUES 
	(?, ?, ?, ?, ?)
`

/*
This function inserts a new comment to the Database

Parameters:

	r types.CommentCreationRequest
*/
func SaveCommentInDB(r types.Comment) error {
	stmt, err := storage.DB_Conn.Prepare(SaveCommentInDBQuery)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if err != nil {
		utils.ErrorConsoleLog("Comment creation fail!")
		return err
	}
	if _, err := stmt.Exec(r.ID, r.Post_ID, r.User.ID.String(), time.Now().Format("YYYY-MM-DD"), r.Content); err != nil {
		return errors.Join(types.ErrExec, err)
	}

	return nil
}

func ConstructNewCommentFromRequest(r types.CommentCreationRequest) (types.Comment, error) {
	new_pid, err := uuid.NewV7()
	if err != nil {
		return (types.Comment{}), errors.Join(types.ErrUUID, err)
	}

	author_session, ok := types.Sessions[uuid.FromStringOrNil(r.Session_id)]
	if !ok {
		return (types.Comment{}), errors.Join(types.ErrSessionNotFound, err)
	}

	comment_author := author_session.User

	return types.Comment{
		ID:           new_pid,
		User:         comment_author,
		Post_ID:      new_pid,
		CreationDate: time.Now(),
		Content:      r.Comment_text,
		Likes:        0,
	}, nil
}

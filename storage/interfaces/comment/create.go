package comment

import (
	"errors"
	"time"

	"RTF/storage"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const SaveCommentInDBQuery = `INSERT INTO comments (post_id, user_id, comment_date, comment) VALUES 
	(?, ?, ?, ?)
`

/*
This function inserts a new comment to the Database

Parameters:
	r types.CommentCreationRequest
*/
func SaveCommentInDB(r types.CommentCreationRequest) error {
	stmt, err := storage.DB_Conn.Prepare(SaveCommentInDBQuery)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	userid := types.Sessions[uuid.FromStringOrNil(r.Session_id)].User.ID

	if _, err := stmt.Exec(r.Post_id, userid.String(), time.Now().Format("YYYY-MM-DD"), r.Comment_text); err != nil {
		return errors.Join(types.ErrExec, err)
	}

	return nil
}

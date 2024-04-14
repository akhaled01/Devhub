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

	postid       uuid.UUID
	userid       uuid.UUID
	comment_date time.Time
	comment      string
*/
func SaveCommentInDB(postid uuid.UUID, userid uuid.UUID, comment_date time.Time, comment string) error {
	stmt, err := storage.DB_Conn.Prepare(SaveCommentInDBQuery)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String(), userid.String(), comment_date.Format("YYYY-MM-DD"), comment); err != nil {
		return errors.Join(types.ErrExec, err)
	}

	return nil
}

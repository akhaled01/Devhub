package likes

import (
	"errors"

	"RTF/storage"

	"github.com/gofrs/uuid"
)

const (
	CREATE_LIKE_RECORD = `INSERT INTO post_likes (post_id, user_id) 
	VALUES (?, ?)`
	DELETE_LIKE_RECORD = `DELETE FROM post_likes WHERE post_id = ? 
	AND user_id = ?
`
)

/*
Invokes a query that inserts a new like into the DB
*/
func CreateLikeRecordPost(postid uuid.UUID, userid uuid.UUID) error {
	stmt, err := storage.DB_Conn.Prepare(CREATE_LIKE_RECORD)
	if err != nil {
		return errors.Join(errors.New("error preparing CreateLikeRecord query"), err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String(), userid.String()); err != nil {
		return errors.Join(errors.New("error executing CreateLikeRecord query"), err)
	}
	return nil
}

/*
Invokes a query that deletes a like from the DB
*/
func DeleteLikeRecordPost(postid uuid.UUID, userid uuid.UUID) error {
	stmt, err := storage.DB_Conn.Prepare(DELETE_LIKE_RECORD)
	if err != nil {
		return errors.Join(errors.New("error preparing DeleteLikeRecord query"), err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String(), userid.String()); err != nil {
		return errors.Join(errors.New("error executing DeleteLikeRecord query"), err)
	}
	return nil
}

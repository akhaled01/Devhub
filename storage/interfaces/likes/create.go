package likes

import (
	"errors"

	"RTF/storage"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const (
	CREATE_LIKE_RECORD = `INSERT INTO post_likes (post_id, user_id) VALUES (?, ?)`
	DELETE_LIKE_RECORD = `DELETE FROM post_likes WHERE post_id = ? AND user_id = ?`
	INCREMENT_LIKES    = `UPDATE posts SET like_count = like_count + 1 WHERE post_id = ?`
	DECREMENT_LIKES    = `UPDATE posts SET like_count = like_count - 1 WHERE post_id = ?`
)

// Invokes a query that inserts a new like into the DB
func CreateLikeRecordPost(postid uuid.UUID, userid uuid.UUID) error {
	stmt, err := storage.DB_Conn.Prepare(CREATE_LIKE_RECORD)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String(), userid.String()); err != nil {
		return errors.Join(types.ErrScan, err)
	}

	// Increment likes
	stmt, err = storage.DB_Conn.Prepare(INCREMENT_LIKES)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String()); err != nil {
		return errors.Join(types.ErrExec, err)
	}

	return nil
}

// Invokes a query that deletes a like from the DB
func DeleteLikeRecordPost(postid uuid.UUID, userid uuid.UUID) error {
	stmt, err := storage.DB_Conn.Prepare(DELETE_LIKE_RECORD)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String(), userid.String()); err != nil {
		return errors.Join(types.ErrExec, err)
	}

	// decrement likes
	stmt, err = storage.DB_Conn.Prepare(DECREMENT_LIKES)
	if err != nil {
		return errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postid.String()); err != nil {
		return errors.Join(types.ErrExec, err)
	}
	return nil
}

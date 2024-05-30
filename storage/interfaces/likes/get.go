package likes

import (
	"errors"

	"RTF/storage"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const (
	IS_LIKED_QUERY_POST = `SELECT EXISTS(SELECT 1 FROM post_likes
		WHERE post_id = ? AND user_id = ?) `
	IS_LIKED_QUERY_COMMENT = `SELECT EXISTS(SELECT 1 FROM comment_likes
			WHERE comment_id = ? AND user_id = ?) `
	POST_LIKES_QUERY = `SELECT count(*) FROM post_likes WHERE post_id = ?`
)

// Takes in a user's id and a post's id and checks if the user liked the post
func CheckUserPostLike(postid uuid.UUID, userid uuid.UUID) (bool, error) {
	stmt, err := storage.DB_Conn.Prepare(IS_LIKED_QUERY_POST)
	if err != nil {
		return false, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	var isLiked bool

	if err := stmt.QueryRow(postid.String(), userid.String()).Scan(&isLiked); err != nil {
		return false, errors.Join(types.ErrExec, err)
	}

	return isLiked, nil
}

// function that checks if the user liked a comment
func CheckUserCommentLike(commentid uuid.UUID, userid uuid.UUID) (bool, error) {
	stmt, err := storage.DB_Conn.Prepare(IS_LIKED_QUERY_COMMENT)
	if err != nil {
		return false, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()
	var isLiked bool

	if err := stmt.QueryRow(commentid.String(), userid.String()).Scan(&isLiked); err != nil {
		return false, errors.Join(types.ErrExec, err)
	}
	return isLiked, nil
}

// function that gets post likes by id.
// call this function to get the number of likes on a post
func GetPostLikes(postID uuid.UUID) (int64, error) {
	row := storage.DB_Conn.QueryRow(POST_LIKES_QUERY, postID.String())
	var count int64
	if err := row.Scan(&count); err != nil {
		return -1, errors.Join(types.ErrScan, err)
	}

	return count, nil
}

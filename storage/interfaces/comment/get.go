package comment

import (
	"database/sql"
	"errors"
	"time"

	"RTF/storage"
	"RTF/storage/interfaces/user"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const QueryPostComments = "SELECT comm_id, user_id, SUBSTR(comment_date, '%Y-%m-%d') AS Date, comment FROM comments WHERE post_id = ?"

// function to return an array of post comments by id
func GetPostCommentsByID(postid uuid.UUID) ([]types.Comment, error) {
	stmt, err := storage.DB_Conn.Prepare(QueryPostComments)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postid.String())
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}

	var comments []types.Comment

	for rows.Next() {
		c := &types.Comment{}
		var comment_id string
		var user_id string
		var comment_date string

		if err := rows.Scan(&comment_id, &user_id, &comment_date, &c.Content); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		u, err := user.GetSingleUser("user_id", user_id)
		if err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}
		partial_user := types.PartialUser{
			ID:       u.ID,
			Username: u.Username,
		}

		//! MIGHT BE BUGGY FROM HERE ON OUT
		if c.Likes, err = GetCommentLikes(uuid.FromStringOrNil(comment_id)); err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}

		c.ID = uuid.FromStringOrNil(comment_id)
		c.Post_ID = postid
		c.User = partial_user
		if c.CreationDate, err = time.Parse("YYYY-MM-DD", comment_date); err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}
		comments = append(comments, *c)
	}

	return comments, nil
}

// gets comment likes by id
func GetCommentLikes(commentID uuid.UUID) (int64, error) {
	query := "SELECT COUNT(*) FROM comment_likes WHERE comment_id = ?"

	row := storage.DB_Conn.QueryRow(query, commentID)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Join(types.ErrGetLikes, err)
	}

	return count, nil
}

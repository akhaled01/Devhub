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

const QueryPostComments = "SELECT comm_id, user_id, comment_date, comment FROM comments WHERE post_id = ?"

// function to return an array of post comments by id
func GetPostCommentsByID(postid uuid.UUID) ([]types.Comment, error) {
	stmt, err := storage.DB_Conn.Prepare(QueryPostComments)
	if err != nil {
		return nil, errors.Join(errors.New("error preparing GetPostCommentsByID query"), err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postid.String())
	if err != nil {
		return nil, errors.Join(errors.New("error executing GetPostCommentsByID query"), err)
	}

	var comments []types.Comment

	for rows.Next() {
		c := &types.Comment{}
		var comment_id string
		var user_id string
		var comment_date string

		if err := rows.Scan(&comment_id, &user_id, &comment_date, &c.Content); err != nil {
			return nil, errors.Join(errors.New("error mapping to comment struct"), err)
		}

		u, err := user.GetSingleUser("user_id", user_id)
		if err != nil {
			return nil, errors.Join(errors.New("error getting comment creator"), err)
		}

		//! MIGHT BE BUGGY FROM HERE ON OUT
		if c.Likes, err = GetCommentLikes(uuid.FromStringOrNil(comment_id)); err != nil {
			return nil, errors.Join(errors.New("error getting comment creator"), err)
		}

		c.ID = uuid.FromStringOrNil(comment_id)
		c.Post_ID = postid
		c.User = *u
		if c.CreationDate, err = time.Parse("YYYY-MM-DD", comment_date); err != nil {
			return nil, errors.Join(errors.New("error getting comment creation date"), err)
		}
		comments = append(comments, *c)
	}

	return comments, nil
}

// gets comment likes by id
func GetCommentLikes(commentID uuid.UUID) (int64, error) {
	query := "SELECT COUNT(*) FROM comment_interactions WHERE comment_id = ?"

	row := storage.DB_Conn.QueryRow(query, commentID)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Join(errors.New("error querying commment likes"), err)
	}

	return count, nil
}
